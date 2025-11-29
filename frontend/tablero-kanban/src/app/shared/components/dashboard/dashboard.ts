import { Component, OnInit } from '@angular/core';
import { Sidebar } from '../sidebar/sidebar';
import { NgComponentOutlet } from '@angular/common';
import { CreateCardLayer } from '../../../features/create-card-layer/create-card-layer';
import { Board } from '../../../features/boards/models/board';
import { BoardComponent } from '../../../features/boards/components/board-component/board-component';

@Component({
  standalone: true,
  selector: 'app-dashboard',
  imports: [
    NgComponentOutlet, 
    Sidebar,
    CreateCardLayer,
    BoardComponent
  ],
  templateUrl: './dashboard.html',
  styleUrl: './dashboard.css',
})
export class Dashboard {
  currentComponent: any = null;
  option: any = null;
  boards: Board[] = [];
  BoardComponent = BoardComponent;

  showComponent(component: any) {
    this.currentComponent = component;
  }

  handleCreatedBoard(board: Board) {
    this.boards.push(board);
    this.currentComponent = null;
    this.option = Board;
  }
  
}
