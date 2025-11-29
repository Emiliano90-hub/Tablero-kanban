import { Component, EventEmitter, Input, OnChanges, OnInit, Output } from '@angular/core';
import { CreateBoardForm } from "../boards/components/create-board-form/create-board-form";
import { Board } from '../boards/models/board';

@Component({
  standalone: true,
  selector: 'app-create-card-layer',
  imports: [CreateBoardForm],
  templateUrl: './create-card-layer.html',
  styleUrl: './create-card-layer.css',
})
export class CreateCardLayer implements OnInit, OnChanges {
  form = null;

  @Input() component!: any;
  @Output() submitted = new EventEmitter<Board>();

  ngOnInit() {
    this.form = this.component;
  }

  ngOnChanges() {
    this.submitted
  }
}
