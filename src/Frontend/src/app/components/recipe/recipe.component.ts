import { Component, OnInit } from '@angular/core';
import { MealDbService } from '../../services/meal-db.service';
import { HttpResponse } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { NewRecipeComponent } from "../new-recipe/new-recipe.component";
import { RouterLink } from '@angular/router';

@Component({
  selector: 'app-recette',
  standalone: true,
  templateUrl: './recipe.component.html',
  styleUrl: './recipe.component.css',
  imports: [FormsModule, CommonModule, NewRecipeComponent, RouterLink]
})

export class RecipeComponent implements OnInit {
  recipe: any;
  recipeId: any;
  responseContent: any;
  instructions: any;
  ingredients: string[] = [];
  ingredientsCount: number = 0;
  measures: string[] = [];

  constructor(private mealDbService: MealDbService) { }
  ngOnInit(): void {
    var urlElements = window.location.href.split('/');
    this.recipeId = urlElements[urlElements.length - 1]
    this.getRecipe();
  }

  getRecipe(): any {
    const observable = this.mealDbService.getRecipeById(this.recipeId);

    // Send request
    observable.subscribe({
      next: (event: any) => {
        if (event instanceof HttpResponse) {
          this.recipe = event.body.meals[0];

          this.instructions = this.recipe.strInstructions.split('\r\n');
          this.instructions = this.instructions.filter((item: string) => item.trim() != '');

          for (let index = 1; index <= 20; index++) {
            var strIngredient = 'strIngredient' + index;
            var strMeasure = 'strMeasure' + index;

            if (this.recipe[strIngredient] != '' && this.recipe[strIngredient] != null) {
              this.ingredients.push(this.recipe[strIngredient]);
              this.ingredientsCount++;
            }

            if(this.recipe[strMeasure] != '' && this.recipe[strMeasure] != null) {
              this.measures.push(this.recipe[strMeasure] + ' ' + this.recipe[strIngredient]);
            }
          }
        }
      },
      error: (httpErrorResponse: any) => {
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
