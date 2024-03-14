import { Component } from '@angular/core';
import {RouterLink, RouterLinkActive} from "@angular/router";

@Component({
  selector: 'app-transactions',
  standalone: true,
  imports: [
    RouterLink,
    RouterLinkActive
  ],
  templateUrl: './wallet-transactions.component.html',
  styleUrl: './wallet-transactions.component.scss'
})
export class WalletTransactionsComponent {

}
