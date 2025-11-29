import { bootstrapApplication } from '@angular/platform-browser';
import { appConfig } from './app/app.config';
import { App } from './app/app';
import { BoardComponent } from './app/features/boards/components/board-component/board-component';
import { Dashboard } from './app/shared/components/dashboard/dashboard';

const routes = [
  { path: '', redirectTo: 'dashboard', pathMatch: 'full' },
  {
    path: 'dashboard',
    component: Dashboard,
    children: [
      { path: 'board', component: BoardComponent },
    ],
  },
];

bootstrapApplication(App, appConfig).catch((err) => console.error(err));
