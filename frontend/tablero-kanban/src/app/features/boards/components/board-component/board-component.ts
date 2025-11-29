import { Component, Input } from '@angular/core';

@Component({
  standalone: true,
  selector: 'app-board',
  imports: [],
  templateUrl: './board-component.html',
  styleUrl: './board-component.css',
})
export class BoardComponent {
  @Input() boardName: string = '';
}