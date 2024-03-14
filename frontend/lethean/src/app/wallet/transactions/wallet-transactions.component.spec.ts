import { ComponentFixture, TestBed } from '@angular/core/testing';

import { WalletTransactionsComponent } from './wallet-transactions.component';

describe('TransactionsComponent', () => {
  let component: WalletTransactionsComponent;
  let fixture: ComponentFixture<WalletTransactionsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [WalletTransactionsComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(WalletTransactionsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
