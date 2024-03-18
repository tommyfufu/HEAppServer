import { Routes } from '@angular/router';
import { PatientSelectionComponent } from './patient-selection/patient-selection.component';
import { PatientProfileComponent } from './patient-profile/patient-profile.component';

export const routes: Routes = [
    { path: 'select-patient', component: PatientSelectionComponent },
    { path: 'patient-profile', component: PatientProfileComponent },
    { path: '', redirectTo: '/select-patient', pathMatch: 'full' },
  ];