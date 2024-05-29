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
    console.log('Selecting patient:', patient);
    this.selectedPatientSource.next(patient);
    this.currentPatientId = patient._id;
    this.selectedPatientIdSource.next(patient._id);
    this.refreshCurrentPatient(); // Refresh data immediately after selection
  }

  selectPatientById(patientId: string): void {
    this.currentPatientId = patientId;
    console.log(this.currentPatientId);
    this.refreshCurrentPatient(); // Always refresh when selecting by ID
  }

  getCurrentPatientId(): string | null {
    return this.currentPatientId;
  }

  refreshCurrentPatient(): void {
    console.log('refreshCurrentPatient');
    if (this.currentPatientId) {
      this.apiService.getPatient({ id: this.currentPatientId }).subscribe({
        next: (updatedPatient) => {
          console.log('Updated patient data:', updatedPatient);
          this.selectedPatientSource.next(updatedPatient); // Update with fresh data
        },
        error: (error) => console.error('Error fetching patient data', error),
      });
      // console.log(this.currentPatientId);
    }
  }
}
