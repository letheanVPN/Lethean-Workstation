import { Routes } from '@angular/router';
import {DashboardComponent} from "./dashboard/dashboard.component";
import {VpnComponent} from "./vpn/vpn.component";
import {WalletComponent} from "./wallet/wallet.component";
import {BlockchainComponent} from "./blockchain/blockchain.component";
import {SetupComponent} from "./setup/setup.component";
import {ReportsComponent} from "./reports/reports.component";
import {WalletTransactionsComponent} from "./wallet/transactions/wallet-transactions.component";
import {WalletSendComponent} from "./wallet/send/wallet-send.component";
import {WalletReceiveComponent} from "./wallet/receive/wallet-receive.component";
import {WalletSettingsComponent} from "./wallet/settings/wallet-settings.component";

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
    path: 'setup',
    component: SetupComponent
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
