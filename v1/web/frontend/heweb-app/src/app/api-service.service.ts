import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { map, catchError } from 'rxjs/operators';
import { Patient, Medication, Message } from './models/patient.model';

@Injectable({
  providedIn: 'root',
})
export class ApiService {
  private baseUrl = 'http://140.113.151.61:8090';
  constructor(private http: HttpClient) {}

  getPatients(): Observable<Patient[]> {
    return this.http
      .get<Patient[]>(`${this.baseUrl}/patients`, {
        headers: { 'Content-Type': 'application/json; charset=UTF-8' },
      })
      .pipe(catchError(this.handleError));
  }

  // getPatientById(patientID: string): Observable<Patient> {
  //   return this.http
  //     .get<Patient>(`${this.baseUrl}/patient?id=${patientID}`)
  //     .pipe(catchError(this.handleError));
  // }

  // Extended method in Angular service for HEApp
  getPatient(params: {
    id?: string;
    email?: string;
    phone?: string;
  }): Observable<Patient> {
    const queryParamString = new URLSearchParams(params).toString();
    return this.http
      .get<Patient>(`${this.baseUrl}/patient?${queryParamString}`)
      .pipe(catchError(this.handleError));
  }

  getMessagesForPatient(patientID: string): Observable<Message[]> {
    return this.http
      .get<Patient>(`${this.baseUrl}/patient/${patientID}`, {
        headers: { 'Content-Type': 'application/json; charset=UTF-8' },
      })
      .pipe(
        map((patient) => patient.messages),
        catchError(this.handleError)
      );
  }
  // update a patient's information
  updatePatient(patientID: string, patientData: Patient): Observable<Patient> {
    return this.http
      .put<Patient>(`${this.baseUrl}/patient/${patientID}`, patientData, {
        headers: { 'Content-Type': 'application/json; charset=UTF-8' },
      })
      .pipe(catchError(this.handleError));
  }

  addMedicationForPatient(
    patientID: string,
    medications: Medication[]
  ): Observable<Patient> {
    const payload = { medication: medications };
    console.log('Sending medication update:', payload);

    return this.http
      .patch<Patient>(
        `${this.baseUrl}/patient/${patientID}/medication`,
        payload,
        { headers: { 'Content-Type': 'application/json; charset=UTF-8' } }
      )
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
