import { Component, OnInit, OnDestroy } from '@angular/core';
import {
  FormBuilder,
  FormGroup,
  FormArray,
  ReactiveFormsModule,
} from '@angular/forms';
import { PatientDataService } from '../shared/patient-data.service';
import { CommonModule } from '@angular/common';
import { Patient, Medication } from '../models/patient.model';
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
export class MedicationEntryComponent implements OnInit, OnDestroy {
  medicationForm: FormGroup;
  private unsubscribe$ = new Subject<void>();

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
          this.initFormWithMedications(patient.medications);
        } else {
          this.clearFormArray(this.medications);
        }
      });
  }

  ngOnDestroy(): void {
    this.unsubscribe$.next();
    this.unsubscribe$.complete();
  }

  get medications(): FormArray {
    return this.medicationForm.get('medications') as FormArray;
  }

  initFormWithMedications(medications: Medication[] | null | undefined): void {
    this.medications.clear();
    (medications || []).forEach(med => {
        this.medications.push(this.fb.group({
            Name: [med.name],
            Dosage: [med.dosage],
            Frequency: [med.frequency]
        }));
    });
}

  addMedicationField(): void {
    this.medications.push(
      this.fb.group({
        Name: [''],
        Dosage: [0],
        Frequency: [0],
      })
    );
  }

  submit(): void {
    const patientID = this.patientDataService.getCurrentPatientId();
    if (!patientID) {
      console.error('No patient selected');
      return;
    }

    const updatedMedications: Medication[] =
      this.medicationForm.value.medications;
    this.apiService
      .addMedicationForPatient(patientID, updatedMedications)
      .subscribe({
        next: (updatedPatient) => {
          console.log('Updated patient medication:', updatedPatient);
          alert('Medication updated successfully!');
          this.patientDataService.refreshCurrentPatient();
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

  private clearFormArray(formArray: FormArray): void {
    while (formArray.length !== 0) {
      formArray.removeAt(0);
    }
  }
}
