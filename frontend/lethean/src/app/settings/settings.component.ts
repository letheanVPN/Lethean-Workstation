import {Component, OnInit} from '@angular/core';
import {RouterLink} from "@angular/router";

@Component({
  selector: 'app-settings',
  standalone: true,
  imports: [
    RouterLink
  ],
  templateUrl: './settings.component.html',
  styleUrl: './settings.component.scss'
})
export class SettingsComponent implements OnInit{
    ngOnInit(): void {
        console.log('SettingsComponent');
        throw new Error('Method not implemented.');
    }

}

