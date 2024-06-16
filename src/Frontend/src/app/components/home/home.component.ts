import { Component, OnInit } from '@angular/core';
import { MealDbService } from '../../services/meal-db.service';
import { DomSanitizer } from '@angular/platform-browser';
import { HttpResponse } from '@angular/common/http';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css'
})
export class HomeComponent implements OnInit {
  platDuJour: any;
  responseContent: any;
  // Meal properties
  meal: any;
  mealImg: any
  category: any;
  area: any;
  tags: any;
  measures: any = [];
  ingredients: any = [];
  ingredientsCount: number = 0;
  instructions: any;
  mealVideo: any;

  constructor(private mealDbService: MealDbService, private sanitizer: DomSanitizer) {

  }
  ngOnInit(): void {
    this.showPlatDuJour();
  }

  showPlatDuJour() {
    // console.log("test");
    var observable = this.mealDbService.getRandomMeal();

    // Send request
    observable.subscribe({
      next: (event: any) => {
        // console.log(event);
        if (event instanceof HttpResponse) {
          this.platDuJour = event.body.meals[0];
          this.mealImg = this.platDuJour.strMealThumb;
          this.meal = this.platDuJour.strMeal;
          this.category = this.platDuJour.strCategory;
          this.tags = this.platDuJour.strTags;
          this.area = this.platDuJour.strArea;
          this.instructions = this.platDuJour.strInstructions.split('\r\n');
          this.instructions = this.instructions.filter((item: string) => item.trim() != '');
          this.mealVideo = "https://youtube.com/embed/" + this.platDuJour.strYoutube.split('=')[1];
          this.mealVideo = this.sanitizer.bypassSecurityTrustResourceUrl(this.mealVideo);
          this.ingredients = [];
          this.measures = [];
          this.ingredientsCount = 0;
          for (let index = 1; index <= 20; index++) {
            var strIngredient = 'strIngredient' + index;
            var strMeasure = 'strMeasure' + index;
            if (this.platDuJour[strIngredient] != '') {
              // console.log(this.platDuJour[strIngredient]); 
              this.ingredients.push(this.platDuJour[strIngredient]);
              this.ingredientsCount++;
              // console.log(this.ingredients);
            }
            if(this.platDuJour[strMeasure] != '') {
              this.measures.push(this.platDuJour[strMeasure] + ' ' + this.platDuJour[strIngredient]);
            }
          }

          // Remove blank elements
          this.measures = this.measures.filter((item: string) => item.trim() != '');
        }
      },
      error: (httpErrorResponse: any) => {
        // console.log(httpErrorResponse);
        if (httpErrorResponse.error) {
          this.responseContent = httpErrorResponse.error;
        } else {
          this.responseContent = 'Une erreur est survenue !';
        }
      },
      complete: () => {
      },
    });
  }
}
