import { Component, Input, OnInit } from '@angular/core';
import { Cat } from 'src/app/cat';
import { CatsService } from '../cats.service';

declare var M: any;
@Component({
  selector: 'app-cat',
  templateUrl: './cat.component.html',
  styleUrls: ['./cat.component.css']
})
export class CatComponent implements OnInit {
  @Input() cat?: Cat;
  @Input() rating?: number;
  constructor(private catsService: CatsService) {
  }

  ngOnInit(): void {
  }

  vote(rating: number) {
    if(this.cat == null) {
      return;
    }
    const result = this.catsService.vote(this.cat, rating);
    result.forEach(() => {
      M.toast({html: `Voto enviado para ${this.cat?.name}!`});
    });
  }

}
