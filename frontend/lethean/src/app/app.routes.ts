import { Routes } from '@angular/router';
import {DashboardComponent} from "./dashboard/dashboard.component";
import {VpnComponent} from "./vpn/vpn.component";
import {WalletComponent} from "./wallet/wallet.component";
import {BlockchainComponent} from "./blockchain/blockchain.component";
import {SettingsComponent} from "./settings/settings.component";
import {ReportsComponent} from "./reports/reports.component";
import {WalletTransactionsComponent} from "./wallet/transactions/wallet-transactions.component";
import {WalletSendComponent} from "./wallet/send/wallet-send.component";
import {WalletReceiveComponent} from "./wallet/receive/wallet-receive.component";
import {WalletSettingsComponent} from "./wallet/settings/wallet-settings.component";
import {SettingsFullNodeComponent} from "./settings/full-node/settings-full-node.component";
import {SettingsDockerComponent} from "./settings/docker/settings-docker.component";

export const routes: Routes = [

  {
    path: 'dashboard',
    component: DashboardComponent,
  },
  {
    path: 'vpn',
    component: VpnComponent
  },
  {
    path: 'blockchain',
    component: BlockchainComponent
  },
  {
    path: 'wallet',
    component: WalletComponent
  },
  {
    path: 'wallet/transactions',
    component: WalletTransactionsComponent
  },
  {
    path: 'wallet/send',
    component: WalletSendComponent
  },
  {
    path: 'wallet/receive',
    component: WalletReceiveComponent
  },
  {
    path: 'wallet/settings',
    component: WalletSettingsComponent
  },
  {
    path: 'settings',
    component: SettingsComponent
  },
  {
    path: 'settings/full-node',
    component: SettingsFullNodeComponent

  },
  {
    path: 'settings/docker',
    component: SettingsDockerComponent

  },
  {
    path: 'reports',
    component: ReportsComponent
  },
  {
    path: '',
    redirectTo: 'dashboard',
    pathMatch: 'full',
  }
];
