import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { PatientDataService } from '../shared/patient-data.service';
import { Message, Patient } from '../models/patient.model';

@Component({
  selector: 'app-patient-profile',
  templateUrl: './patient-profile.component.html',
  styleUrls: ['./patient-profile.component.css'],
  standalone: true,
  imports: [CommonModule],
})
export class PatientProfileComponent implements OnInit {
  selectedPatient: Patient | null = null;

  constructor(private patientDataService: PatientDataService) {}

  ngOnInit(): void {
    this.patientDataService.selectedPatient$.subscribe((patient) => {
      console.log('Received patient data:', patient);
      this.selectedPatient = patient;
    });
  }

  getMessages(): Message[] {
    return this.selectedPatient?.messages || [];
  }
}
