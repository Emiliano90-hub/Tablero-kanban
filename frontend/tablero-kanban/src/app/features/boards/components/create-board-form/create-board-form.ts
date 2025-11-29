import { Component, EventEmitter, Output } from '@angular/core';
import { FormBuilder, FormGroup, Validators, ReactiveFormsModule } from '@angular/forms';
import { BoardService } from '../../services/board-service';
import { Board } from '../../models/board';


@Component({
  standalone: true,
  selector: 'app-create-board-form',
  imports: [ReactiveFormsModule],
  templateUrl: './create-board-form.html',
  styleUrl: './create-board-form.css',
})
export class CreateBoardForm {
  form: FormGroup;
  @Output() createdBoard = new EventEmitter<Board>();

  constructor(private fb: FormBuilder, private boardService: BoardService) {
    this.form = this.fb.group({
      name: ['', Validators.required],
    });
  }

  submit() {
    if (this.form.invalid) return;
    const newBoard: Board = { name: this.form.value.name };
    this.boardService.addBoard(this.form.value.name);
    this.createdBoard.emit(newBoard);
    this.form.reset();
  }
}
