import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SettingsDockerComponent } from './settings-docker.component';

describe('SettingsDockerComponent', () => {
  let component: SettingsDockerComponent;
  let fixture: ComponentFixture<SettingsDockerComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SettingsDockerComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(SettingsDockerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
