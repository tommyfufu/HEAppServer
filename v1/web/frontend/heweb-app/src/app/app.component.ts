import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { PatientSelectionComponent } from './patient-selection/patient-selection.component';
import { PatientProfileComponent } from './patient-profile/patient-profile.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, PatientSelectionComponent, PatientProfileComponent],
  // imports: [RouterOutlet,],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css'
})
export class AppComponent {
  title = 'heweb-app';
}
