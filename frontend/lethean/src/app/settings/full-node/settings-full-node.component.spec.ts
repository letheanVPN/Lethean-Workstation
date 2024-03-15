import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SettingsFullNodeComponent } from './settings-full-node.component';

describe('SetupFullNodeComponent', () => {
  let component: SettingsFullNodeComponent;
  let fixture: ComponentFixture<SettingsFullNodeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SettingsFullNodeComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SettingsFullNodeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
