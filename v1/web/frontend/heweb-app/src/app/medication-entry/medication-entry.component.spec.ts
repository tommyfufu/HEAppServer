import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MedicationEntryComponent } from './medication-entry.component';

describe('MedicationEntryComponent', () => {
  let component: MedicationEntryComponent;
  let fixture: ComponentFixture<MedicationEntryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [MedicationEntryComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(MedicationEntryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
