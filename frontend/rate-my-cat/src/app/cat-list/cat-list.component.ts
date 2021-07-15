import { Component, OnInit } from '@angular/core';
import { Cat } from 'src/app/cat';
import { CatsService } from '../cats.service';

@Component({
  selector: 'app-cat-list',
  templateUrl: './cat-list.component.html',
  styleUrls: ['./cat-list.component.css']
})
export class CatListComponent implements OnInit {
  cats: Cat[] = [];
  constructor(private catsService: CatsService) {

  }
  fetchId?: any;
  ngOnInit(): void {
    this.fetchId = setInterval(() => {
      this.fetchCats();
    }, 500);
  }
  ngOnDestroy(): void {
    if (this.fetchId) {
      clearInterval(this.fetchId);
    }
  }

  fetchCats(): void {
    const self = this
    this.catsService.getCats().forEach((cats: Cat[]) => {
      self.cats = cats;
    });
  }
  trackByFn(index:number, item:Cat) {
    return item.id
 }
}
