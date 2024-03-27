import { Component, OnInit } from '@angular/core';
import {
  FormBuilder,
  FormGroup,
  FormArray,
  ReactiveFormsModule,
} from '@angular/forms';
import { PatientDataService } from '../shared/patient-data.service';
import { CommonModule } from '@angular/common';
import { Patient } from '../models/patient.model';
import { ApiService } from '../api-service.service';
import { Subject } from 'rxjs';
import { takeUntil } from 'rxjs/operators';

@Component({
  selector: 'app-medication-entry',
  templateUrl: './medication-entry.component.html',
  styleUrls: ['./medication-entry.component.css'],
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule],
})
export class MedicationEntryComponent implements OnInit {
  medicationForm: FormGroup;
  private unsubscribe$ = new Subject<void>();
  private currentPatientId: string | null = null;
  // Dependency injection:
  // - PatientDataService (get medication data of the patient)
  // - ApiService (get patient's data)
  // Initialize the form array
  constructor(
    private fb: FormBuilder,
    private patientDataService: PatientDataService,
    private apiService: ApiService
  ) {
    this.medicationForm = this.fb.group({
      medications: this.fb.array([]),
    });
  }

  ngOnInit(): void {
    this.patientDataService.selectedPatient$
      .pipe(takeUntil(this.unsubscribe$))
      .subscribe((patient) => {
        if (patient) {
          this.initFormWithMedications(patient.Medication);
        } else {
          this.clearFormArray(this.medications);
        }
      });
  }

  ngOnDestroy() {
    this.unsubscribe$.next();
    this.unsubscribe$.complete();
  }

  get medications(): FormArray {
    return this.medicationForm.get('medications') as FormArray;
  }

  initFormWithMedications(medications: string[]) {
    this.medications.clear();
    // Fill the form array with existing medication records
    medications.forEach((medication) => {
      this.medications.push(this.fb.control(medication));
    });
  }

  clearFormArray(formArray: FormArray) {
    while (formArray.length !== 0) {
      formArray.removeAt(0);
    }
  }

  // Add a new empty medication field
  addMedicationField() {
    this.medications.push(this.fb.control(''));
  }

  submit() {
    const patientID = this.patientDataService.getCurrentPatientId();

    if (!patientID) {
      console.error('No patient selected');
      return;
    }

    const updatedMedications = this.medicationForm.value.medications;
    this.apiService
      .addMedicationForPatient(patientID, updatedMedications)
      .subscribe({
        next: (updatedPatient) => {
          console.log('Updated patient medication:', updatedPatient);
          alert('Medication updated successfully!');
        },
        error: (error) => {
          console.error('Error updating medication:', error);
          alert('Failed to update medication. Please try again later.');
        },
      });
  }
  removeMedicationField(index: number): void {
    this.medications.removeAt(index);
  }
}
