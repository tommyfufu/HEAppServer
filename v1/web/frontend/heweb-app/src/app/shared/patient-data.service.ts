import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable } from 'rxjs';
import { Patient } from '../models/patient.model'; // Adjust the path as needed

@Injectable({
  providedIn: 'root'
})
export class PatientDataService {
  private selectedPatientSource = new BehaviorSubject<Patient | null>(null);
  selectedPatient$ = this.selectedPatientSource.asObservable();

  constructor() { }

  selectPatient(patient: Patient): void {
    this.selectedPatientSource.next(patient);
  }
}
