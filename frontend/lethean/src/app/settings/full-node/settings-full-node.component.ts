import {Component, OnInit} from '@angular/core';
import {NgIf} from "@angular/common";

@Component({
  selector: 'app-settings-full-node',
  standalone: true,
  imports: [
    NgIf
  ],
  templateUrl: './settings-full-node.component.html',
  styleUrl: './settings-full-node.component.scss'
})
export class SettingsFullNodeComponent implements OnInit{

    status_docker: string = 'Checking...';
    status_lethean_node: string = 'Checking...';
    status_lethean_image: string = 'Checking...';
    ngOnInit(): void {

      const docker = fetch('http://localhost:36911/docker/container/list')
        .then(response => response.json()).then(data => {
          if(data.length >= 0){
            this.status_docker = 'Installed'
            for (let i = 0; i < data.length; i++) {
              if(data[i].Names.includes('/letheannode')){
                this.status_lethean_node = 'Installed'
              }
            }
            if(this.status_lethean_node != 'Installed'){
              this.status_lethean_node = 'Not Installed'
            }
            fetch('http://localhost:36911/docker/image/list').then(response => response.json()).then(data => {
              for (let i = 0; i < data.length; i++) {
                if (data[0].RepoTags.includes('limosek/lvpn:dev')) {
                  this.status_lethean_image = 'Installed'
                }
              }

              if(this.status_lethean_image != 'Installed'){
                this.status_lethean_image = 'Not Installed'
              }
            })
          }
        });
        throw new Error('Method not implemented.');
    }

  pullImage() {
    fetch('http://localhost:36911/docker/image/pull', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({image: 'limosek/lvpn:dev'})
    }).then(response => response.json()).then(data => {
      console.log(data)
    })
  }

  createContainer() {

      const container = {
        Image: "limosek/lvpn:dev",
        Volumes: {
          "/home/lvpn": {}
        },
        HostConfig: {
          "CapAdd": [
            "NET_ADMIN"
          ],
          "PortBindings": {
            "8880/tcp": [
              {
                "HostPort": "8880"
              }
            ],
            "8880/udp": [
              {
                "HostPort": "8880"
              }
            ],
            "8881/tdp": [
              {
                "HostPort": "8881"
              }
            ],
            "8881/udp": [
              {
                "HostPort": "8881"
              }
            ]
          }
        },
        "ExposedPorts": {
          "8880/tcp": {},
          "8880/udp": {},
          "8881/tdp": {},
          "8881/udp": {}
        },
        }


    fetch('http://localhost:36911/docker/container/create/letheannode', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(container)
    }).then(response => response.json()).then(data => {
      console.log(data)
    })
  }
}
