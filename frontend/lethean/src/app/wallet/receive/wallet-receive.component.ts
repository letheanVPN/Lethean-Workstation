import { Component } from '@angular/core';
import {RouterLink, RouterLinkActive} from "@angular/router";

@Component({
  selector: 'app-receive',
  standalone: true,
  imports: [
    RouterLink,
    RouterLinkActive
  ],
  templateUrl: './wallet-receive.component.html',
  styleUrl: './wallet-receive.component.scss'
})
export class WalletReceiveComponent {

}
