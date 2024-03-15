import {Component, OnInit} from '@angular/core';
import {NgForOf, NgIf} from "@angular/common";

@Component({
  selector: 'app-docker',
  standalone: true,
  imports: [
    NgIf,
    NgForOf
  ],
  templateUrl: './settings-docker.component.html',
  styleUrl: './settings-docker.component.scss'
})
export class SettingsDockerComponent implements OnInit {
  containers: any;

  ngOnInit(): void {

    this.getContainers()
  }

  getContainers() {
    fetch('http://localhost:36911/docker/container/list')
      .then(response => response.json()).then(data => {
        this.containers = data;
      });
  }

  deleteContainer(Id: string) {
    fetch(`http://localhost:36911/docker/container/remove/${Id}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    }).then(response => response.json()).then(data => {
      this.getContainers()
    })
  }

  stopContainer(Id: string) {
    fetch(`http://localhost:36911/docker/container/stop/${Id}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    }).then(response => response.json()).then(data => {
      this.getContainers()
    })
  }

  startContainer(Id: string) {
    fetch(`http://localhost:36911/docker/container/start/${Id}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    }).then(response => response.json()).then(data => {
      this.getContainers()
    })
  }
}
