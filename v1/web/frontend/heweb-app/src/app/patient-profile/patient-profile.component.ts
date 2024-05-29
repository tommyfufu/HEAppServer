import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { PatientDataService } from '../shared/patient-data.service';
import { Patient, Message } from '../models/patient.model';
import { ApiService } from '../api-service.service';

@Component({
  selector: 'app-patient-profile',
  templateUrl: './patient-profile.component.html',
  styleUrls: ['./patient-profile.component.css'],
  standalone: true,
  imports: [CommonModule, FormsModule], // Ensure FormsModule is imported for ngModel to work
})
export class PatientProfileComponent implements OnInit {
  selectedPatient: Patient | null = null;
  tempPatient: Patient | null = null; // Temporary patient object for binding form fields in edit mode
  editMode: boolean = false; // Flag to toggle edit mode

  constructor(
    private patientDataService: PatientDataService,
    private apiService: ApiService
  ) {}

  ngOnInit(): void {
    this.patientDataService.selectedPatient$.subscribe((patient) => {
      this.selectedPatient = patient;
      this.tempPatient = patient ? { ...patient } : null; // Clone the patient for editing or set to null if patient is undefined
      this.editMode = false; // Ensure edit mode is off when a new patient is selected
    });
  }

  getMessages(): Message[] {
    // Check if there are messages, if not, return an empty array
    if (this.selectedPatient && this.selectedPatient.messages) {
      return this.selectedPatient.messages.slice().reverse();
    } else {
      return [];
    }
  }

  toggleEditMode(): void {
    this.editMode = !this.editMode;
    if (!this.editMode && this.tempPatient) {
      // Save changes or revert them based on user actions
      this.saveChanges();
    } else {
      this.tempPatient = this.selectedPatient
        ? { ...this.selectedPatient }
        : null;
    }
  }

  saveChanges(): void {
    if (this.tempPatient) {
      this.apiService
        .updatePatient(this.tempPatient._id, this.tempPatient)
        .subscribe({
          next: (updatedPatient) => {
            this.selectedPatient = { ...updatedPatient };
            this.editMode = false; // Turn off edit mode on successful save
            this.patientDataService.selectPatient(updatedPatient); // Optionally update globally selected patient
          },
          error: (error) => {
            console.error('Failed to update patient:', error);
            alert('Unable to save changes. Please try again.');
          },
        });
    }
  }

  cancelEdit(): void {
    this.tempPatient = this.selectedPatient
      ? { ...this.selectedPatient }
      : null;
    this.editMode = false; // Ensure edit mode is turned off
  }
}
