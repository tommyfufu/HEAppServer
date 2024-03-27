import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable } from 'rxjs';
import { Patient } from '../models/patient.model'; 

@Injectable({
  providedIn: 'root',
})
export class PatientDataService {
  private selectedPatientSource = new BehaviorSubject<Patient | null>(null);
  selectedPatient$ = this.selectedPatientSource.asObservable();
  private currentPatientId: string | null = null;

  private selectedPatientIdSource = new BehaviorSubject<string | null>(null);
  selectedPatientId$ = this.selectedPatientIdSource.asObservable();

  constructor() {}

  selectPatient(patient: Patient): void {
    this.selectedPatientSource.next(patient);
    this.currentPatientId = patient.ID;
  }

  selectPatientById(patientId: string): void {
    this.selectedPatientIdSource.next(patientId);
  }

  getCurrentPatientId(): string | null {
    return this.selectedPatientIdSource.value;
  }
}
