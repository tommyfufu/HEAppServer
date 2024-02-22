import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common'; 
import { PatientDataService } from '../shared/patient-data.service';
import { Patient } from '../models/patient.model';

@Component({
  selector: 'app-patient-profile',
  templateUrl: './patient-profile.component.html',
  styleUrls: ['./patient-profile.component.css'],
  standalone: true,
  imports: [CommonModule,],
})
export class PatientProfileComponent implements OnInit {
  selectedPatient: Patient | null = null;

  constructor(private patientDataService: PatientDataService) {}

  ngOnInit(): void {
    this.patientDataService.selectedPatient$.subscribe(patient => {
      this.selectedPatient = patient;
    });
  }

  getMessagesArray(messages: Record<string, string>): {date: string, content: string}[] {
    return Object.entries(messages).map(([date, content]) => ({date, content}));
  }
}
