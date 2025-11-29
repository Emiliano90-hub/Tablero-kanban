import { Injectable, signal } from '@angular/core';
import { Board } from '../models/board';

@Injectable({
  providedIn: 'root',
})
export class BoardService {
  boards = signal<Board[]>([]);

  addBoard(name: string) {
    const newBoard: Board = { name };
    this.boards.update((current) => [...current, newBoard]);
  }

  getBoards() {
    return this.boards();
  }
  
}
