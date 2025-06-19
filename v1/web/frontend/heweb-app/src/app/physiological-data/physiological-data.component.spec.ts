import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PhysiologicalDataComponent } from './physiological-data.component';

describe('PhysiologicalDataComponent', () => {
  let component: PhysiologicalDataComponent;
  let fixture: ComponentFixture<PhysiologicalDataComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PhysiologicalDataComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(PhysiologicalDataComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
