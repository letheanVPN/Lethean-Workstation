import { ComponentFixture, TestBed } from '@angular/core/testing';

import { WalletReceiveComponent } from './wallet-receive.component';

describe('WalletReceiveComponent', () => {
  let component: WalletReceiveComponent;
  let fixture: ComponentFixture<WalletReceiveComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [WalletReceiveComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(WalletReceiveComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
