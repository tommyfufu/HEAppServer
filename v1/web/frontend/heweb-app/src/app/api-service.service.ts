import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { map, catchError } from 'rxjs/operators';
import { Patient } from './models/patient.model';

@Injectable({
  providedIn: 'root',
})
export class ApiService {
  private baseUrl = 'http://localhost:8090';
  constructor(private http: HttpClient) {}

  getPatients(): Observable<Patient[]> {
    return this.http
      .get<Patient[]>(`${this.baseUrl}/patients`)
      .pipe(catchError(this.handleError));
  }

  getPatient(patientID: string): Observable<Patient> {
    return this.http
      .get<Patient>(`${this.baseUrl}/patient/${patientID}`)
      .pipe(catchError(this.handleError));
  }
  
  getMessagesForPatient(patientID: string): Observable<Record<string, string>> {
    return this.http.get<Patient>(`${this.baseUrl}/patient/${patientID}`).pipe(
      map((patient) => patient.Messages),
      catchError(this.handleError)
    );
  }
  // Method to update a patient's information
  updatePatient(patientID: string, patientData: Patient): Observable<Patient> {
    return this.http.put<Patient>(`${this.baseUrl}/patient/${patientID}`, patientData)
      .pipe(catchError(this.handleError));
  }

  // Method to add a new medication entry for a patient
  addMedicationForPatient(patientID: string, medication: string[]): Observable<Patient> {
    // Assuming the backend expects the medication array under a "medication" property
    const updatePayload = { medication };
    return this.http.patch<Patient>(`${this.baseUrl}/patient/${patientID}/medication`, updatePayload)
      .pipe(catchError(this.handleError));
  }

  private handleError(error: HttpErrorResponse) {
    console.error(
      `Backend returned code ${error.status}, body was: `,
      error.error
    );
    return throwError('Something bad happened; please try again later.');
  }
}
