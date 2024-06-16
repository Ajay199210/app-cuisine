import { HttpClient, HttpHeaders, HttpRequest } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class RecipesService {
  // private baseUrl = 'http://localhost:8282'
  private baseUrl = "https://14oi3yhq0g.execute-api.ca-central-1.amazonaws.com"

  constructor(private http: HttpClient) { }

  createRecipe(userId: number, recipeName: string, recipeCategory: string,
    recipeInstructions: string, recipeIngredients: string[], recipeMeasures: string[]): Observable<any> {

    const headers = new HttpHeaders();
    headers.set('Content-Type', 'application/json');

    const body = {
      "userId": userId,
      "strMeal": recipeName,
      "strCategory": recipeCategory,
      "strInstructions": recipeInstructions,
      "strIngredient1": recipeIngredients[0],
      "strIngredient2": recipeIngredients[1],
      "strIngredient3": recipeIngredients[2],
      "strIngredient4": recipeIngredients[3],
      "strIngredient5": recipeIngredients[4],
      "strIngredient6": recipeIngredients[5],
      "strIngredient7": recipeIngredients[6],
      "strIngredient8": recipeIngredients[7],
      "strIngredient9": recipeIngredients[8],
      "strIngredient10": recipeIngredients[9],
      "strIngredient11": recipeIngredients[10],
      "strIngredient12": recipeIngredients[11],
      "strIngredient13": recipeIngredients[12],
      "strIngredient14": recipeIngredients[13],
      "strIngredient15": recipeIngredients[14],
      "strIngredient16": recipeIngredients[15],
      "strIngredient17": recipeIngredients[16],
      "strIngredient18": recipeIngredients[17],
      "strIngredient19": recipeIngredients[18],
      "strIngredient20": recipeIngredients[19],
      "strMeasue1": recipeMeasures[0],
      "strMeasue2": recipeMeasures[1],
      "strMeasue3": recipeMeasures[2],
      "strMeasue4": recipeMeasures[3],
      "strMeasue5": recipeMeasures[4],
      "strMeasue6": recipeMeasures[5],
      "strMeasue7": recipeMeasures[6],
      "strMeasue8": recipeMeasures[7],
      "strMeasue9": recipeMeasures[8],
      "strMeasure10": recipeMeasures[9],
      "strMeasure11": recipeMeasures[10],
      "strMeasure12": recipeMeasures[11],
      "strMeasure13": recipeMeasures[12],
      "strMeasure14": recipeMeasures[13],
      "strMeasure15": recipeMeasures[14],
      "strMeasure16": recipeMeasures[15],
      "strMeasure17": recipeMeasures[16],
      "strMeasure18": recipeMeasures[17],
      "strMeasure19": recipeMeasures[18],
      "strMeasure20": recipeMeasures[19],
    };

    const postReq = new HttpRequest("POST", `${this.baseUrl}/userRecipes`, body, { headers });

    return this.http.request(postReq);
  }
}
