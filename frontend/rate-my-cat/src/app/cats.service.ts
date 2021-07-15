import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Cat } from 'src/app/cat';
import { BehaviorSubject, Observable, throwError, of } from 'rxjs';
import { map } from 'rxjs/operators';
import { environment } from 'src/environments/environment';

interface CatResponse {
  ID: number;
  Name: string;
  PhotoURL: string;
  Rating?: number;
}

interface ListCatsResponse {
  Cats: CatResponse[];
}

@Injectable({
  providedIn: 'root'
})
export class CatsService {

  constructor(private http: HttpClient) { }

  getCats(): Observable<Cat[]> {
    return this
      .http.get<ListCatsResponse>(environment.backendURL + "/listCats")
      .pipe(map((response: ListCatsResponse):Cat[] =>{
        return response.Cats.map((catResponse: CatResponse): Cat => ({
          id: catResponse.ID,
          name: catResponse.Name,
          photoURL: catResponse.PhotoURL,
          rating: catResponse.Rating,
        }))
      }))
  }

  vote(cat: Cat, rating: number): Observable<Object> {
    const payload = {
      CatID: cat.id,
      Vote: rating
    };
    return this.http.post(environment.backendURL +"/vote", payload);

  }
}
