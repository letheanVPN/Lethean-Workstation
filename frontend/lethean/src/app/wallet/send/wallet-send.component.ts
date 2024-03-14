import { Component } from '@angular/core';
import {RouterLink, RouterLinkActive} from "@angular/router";

@Component({
  selector: 'app-send',
  standalone: true,
  imports: [
    RouterLink,
    RouterLinkActive
  ],
  templateUrl: './wallet-send.component.html',
  styleUrl: './wallet-send.component.scss'
})
export class WalletSendComponent {

}
