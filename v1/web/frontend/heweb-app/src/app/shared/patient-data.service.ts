import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable } from 'rxjs';
import { Patient } from '../models/patient.model';
import { ApiService } from '../api-service.service';

@Injectable({
  providedIn: 'root',
})
export class PatientDataService {
  private selectedPatientSource = new BehaviorSubject<Patient | null>(null);
  selectedPatient$ = this.selectedPatientSource.asObservable();
  private currentPatientId: string | null = null;

  private selectedPatientIdSource = new BehaviorSubject<string | null>(null);
  selectedPatientId$ = this.selectedPatientIdSource.asObservable();

  constructor(private apiService: ApiService) {}

  selectPatient(patient: Patient): void {
    this.selectedPatientSource.next(patient);
    this.currentPatientId = patient.ID;
    this.selectedPatientIdSource.next(patient.ID);
  }

  selectPatientById(patientId: string): void {
    this.selectedPatientIdSource.next(patientId);
    this.currentPatientId = patientId; 
    this.refreshCurrentPatient(); 
  }

  getCurrentPatientId(): string | null {
    return this.currentPatientId;
  }

  refreshCurrentPatient(): void {
    if (this.currentPatientId) {
      this.apiService
        .getPatient(this.currentPatientId)
        .subscribe((updatedPatient) => {
          this.selectPatient(updatedPatient);
        });
    }
  }
}
