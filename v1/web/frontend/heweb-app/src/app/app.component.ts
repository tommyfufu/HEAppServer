import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { PatientSelectionComponent } from './patient-selection/patient-selection.component';
import { PatientProfileComponent } from './patient-profile/patient-profile.component';
import { MedicationEntryComponent } from './medication-entry/medication-entry.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, PatientSelectionComponent, PatientProfileComponent, MedicationEntryComponent],
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css',]
})
export class AppComponent {
  title = 'heweb-app';
}
