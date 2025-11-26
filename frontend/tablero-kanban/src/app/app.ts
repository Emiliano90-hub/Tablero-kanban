import { Component, signal } from '@angular/core';
import { Navbar } from './shared/components/navbar/navbar';
import { Sidebar } from './shared/components/sidebar/sidebar';
import { Dashboard } from './shared/components/dashboard/dashboard';
import { Footer } from './shared/components/footer/footer';

@Component({
  standalone: true,
  selector: 'app-root',
  imports: [
    Navbar,
    Sidebar,
    Dashboard,
    Footer
  ],
  templateUrl: './app.html',
  styleUrl: './app.css'
})
export class App {
  protected readonly title = signal('tablero-kanban');
}
