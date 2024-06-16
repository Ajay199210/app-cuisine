import { Component, OnInit } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { RecipeComponent } from "./components/recipe/recipe.component";
import { HeaderComponent } from './components/header/header.component';
import { FooterComponent } from './components/footer/footer.component';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { HomeComponent } from "./components/home/home.component";
import { StorageService } from './services/storage.service';

@Component({
  selector: 'app-root',
  standalone: true,
  templateUrl: './app.component.html',
  styleUrl: './app.component.css',
  imports: [
    RouterOutlet, HeaderComponent, FooterComponent,
    RecipeComponent, FontAwesomeModule,
    HomeComponent
  ]
})
export class AppComponent {
  title = 'recette-remix';

  constructor() { }
}
