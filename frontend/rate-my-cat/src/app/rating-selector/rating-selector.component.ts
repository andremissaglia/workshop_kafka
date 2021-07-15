import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import {FormControl, FormGroup} from '@angular/forms';

@Component({
  selector: 'app-rating-selector',
  templateUrl: './rating-selector.component.html',
  styleUrls: ['./rating-selector.component.css']
})
export class RatingSelectorComponent implements OnInit {
  form = new FormGroup({
    rating: new FormControl(),
  });

  @Output() voteEvent = new EventEmitter<number>();

  constructor() { }

  ngOnInit(): void {
  }
  vote() {
    const rating = parseInt(this.form.value['rating']);
    this.voteEvent.emit(rating);
    this.form.reset();
  }
}
