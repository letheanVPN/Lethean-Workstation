import { Component } from '@angular/core';
import {RouterLink, RouterLinkActive} from "@angular/router";

@Component({
  selector: 'app-settings',
  standalone: true,
  imports: [
    RouterLink,
    RouterLinkActive
  ],
  templateUrl: './wallet-settings.component.html',
  styleUrl: './wallet-settings.component.scss'
})
export class WalletSettingsComponent {

}
