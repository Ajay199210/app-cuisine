import { Component } from '@angular/core';
import { RecipesService } from '../../services/recipes.service';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { RecipeComponent } from '../recipe/recipe.component';

@Component({
  selector: 'app-new-recipe',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './new-recipe.component.html',
  styleUrl: './new-recipe.component.css'
})
export class NewRecipeComponent {
  constructor(private recipesService: RecipesService) { }

  userId?: number;
  newRecipeName: string = '';
  newRecipeCategory: string = '';
  newRecipeInstructions: string = '';
  newRecipeIngredients: string[] = [];
  newRecipeMeasures: string[] = [];
  responseContent: string = '';
  recipeIngredientsCount: number = 2;

  // Add recipe
  addRecipe(): any {
    var observable = this.recipesService.createRecipe(123, this.newRecipeName,
      this.newRecipeCategory, this.newRecipeInstructions, this.newRecipeIngredients, this.newRecipeMeasures);

    // Process ingredients
    this.newRecipeIngredients = [];
    var ingredients = document.getElementsByClassName('ingredient');
    for (let i = 0; i < ingredients.length; i++) {
      var ingredient = ingredients[i] as HTMLInputElement;
      this.newRecipeIngredients.push(ingredient.value);
    }
    console.log(this.newRecipeIngredients);

    // Send request
    // observable.subscribe({
    //   next: (event: any) => {
    //     if (event instanceof HttpResponse) {
    //       console.log(event);
    //       // this.responseContent = event.body.message;
    //     }
    //   },
    //   error: (httpErrorResponse: any) => {
    //     // console.log(httpErrorResponse);
    //     if (httpErrorResponse.error) {
    //       this.responseContent = httpErrorResponse.error;
    //     } else {
    //       this.responseContent = 'Une erreur est survenue !';
    //     }
    //     console.log(this.responseContent);
    //   },
    //   complete: () => {
    //   },
    // });
  }

  // Add ingredient & measure input elements
  addInputs(event: Event): void {
    event.preventDefault();

    this.recipeIngredientsCount++

    var newInputs = document.getElementById("newInputs") as HTMLElement;

    var rowDiv = document.createElement('div');
    rowDiv.classList.add('row', 'my-2');

    var col1Div = document.createElement('div');
    col1Div.classList.add('col');

    // Ingredient input
    var ingredientInput = document.createElement('input');
    ingredientInput.type = 'text';
    ingredientInput.classList.add('form-control', 'ingredient');
    ingredientInput.placeholder = 'Ingredient ' + this.recipeIngredientsCount;

    var validFeedback1 = document.createElement('div');
    validFeedback1.classList.add('valid-feedback');
    validFeedback1.textContent = 'Looks good!';

    col1Div.appendChild(ingredientInput);
    col1Div.appendChild(validFeedback1);

    // Measure input
    var col2Div = document.createElement('div');
    col2Div.classList.add('col');

    var measureInput = document.createElement('input');
    measureInput.type = 'text';
    measureInput.classList.add('form-control', 'measure');
    measureInput.placeholder = 'Measure ' + this.recipeIngredientsCount;

    var validFeedback2 = document.createElement('div');
    validFeedback2.classList.add('valid-feedback');
    validFeedback2.textContent = 'Looks good!';

    // Append new input elements to the newly created block
    col2Div.appendChild(measureInput);
    col2Div.appendChild(validFeedback2);

    rowDiv.appendChild(col1Div);
    rowDiv.appendChild(col2Div);

    newInputs.appendChild(rowDiv);
  }
}
