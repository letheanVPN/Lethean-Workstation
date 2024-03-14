import { ComponentFixture, TestBed } from '@angular/core/testing';

import { WalletSendComponent } from './wallet-send.component';

describe('WalletSendComponent', () => {
  let component: WalletSendComponent;
  let fixture: ComponentFixture<WalletSendComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [WalletSendComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(WalletSendComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
