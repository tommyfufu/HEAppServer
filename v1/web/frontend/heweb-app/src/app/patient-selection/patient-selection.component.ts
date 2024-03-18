import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ApiService } from '../api-service.service';
import { HttpClientModule } from '@angular/common/http';
import { Patient } from '../models/patient.model';
import { PatientDataService } from '../shared/patient-data.service';
@Component({
  selector: 'app-patient-selection',
  standalone: true,
  imports: [CommonModule, HttpClientModule],
  templateUrl: './patient-selection.component.html',
  styleUrls: ['./patient-selection.component.css'],
})
export class PatientSelectionComponent implements OnInit {
  patients: Patient[] = [];
  selectedPatientId: string | null = null; // Assuming patients have an ID

  constructor(
    private apiService: ApiService,
    private patientDataService: PatientDataService
  ) {}

  ngOnInit(): void {
    this.fetchPatients();
  }

  fetchPatients(): void {
    this.apiService.getPatients().subscribe({
      next: (data: Patient[]) => {
        this.patients = data;
        console.log('Fetched patients:', this.patients);
      },
      error: (error) => {
        console.error('Error fetching patients:', error);
      },
      complete: () => console.log('Fetch patients completed'),
    });
  }

  onPatientSelect(event: Event): void {
    const target = event.target as HTMLSelectElement;
    const patientId = target.value;

    const selectedPatient = this.patients.find(
      (patient) => patient.ID === patientId
    );
    if (selectedPatient) {
      this.patientDataService.selectPatient(selectedPatient);
    }
  }
}
