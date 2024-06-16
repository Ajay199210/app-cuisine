import { Component, Output, EventEmitter, OnInit } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { UsersService } from '../../services/users.service';
import { HttpResponse } from '@angular/common/http';
import { StorageService } from '../../services/storage.service';
import { RouterLink } from '@angular/router';

@Component({
  selector: 'app-auth',
  standalone: true,
  imports: [FormsModule, CommonModule, RouterLink],
  templateUrl: './auth.component.html',
  styleUrl: './auth.component.css'
})
export class AuthComponent implements OnInit {
  username: string = '';
  email: string = '';
  password: string = '';
  confirmPass: string = '';
  responseContent: any = '';
  isUserAuthenticated: boolean = false;
  isLoginFailed: boolean = false;
  isSpinnerToggled: boolean = false;

  // Send data (if the user is authenticated) from child to parent
  @Output() userAuthEvent = new EventEmitter<boolean>();

  constructor(private userService: UsersService, private storageService: StorageService) { }

  ngOnInit(): void {
    this.isUserAuthenticated = this.storageService.isLoggedIn();

    if (this.isUserAuthenticated) {
      this.username = this.storageService.getUsername();
    }
  }

  // Login user
  login(): any {
    // Show spinner
    this.isSpinnerToggled = this.toggleSpinner(true, 0);

    // var observable = this.userService.loginUser("test", "TEST123abc!");
    var observable = this.userService.loginUser(this.username, this.password);

    // Send request
    observable.subscribe({
      next: (event: any) => {
        if (event instanceof HttpResponse) {
          this.responseContent = event.body;

          // If we have a JWT token, it means user authenticated
          if (this.responseContent.token) {
            console.log(this.responseContent);

            this.storageService.saveUser(this.username, this.responseContent.token);
            this.isUserAuthenticated = true;

            // Hide sign-in modal
            var modalCloseBtn = document.getElementsByClassName('btnClose')[0] as HTMLButtonElement;
            modalCloseBtn.click();

            window.location.reload();
          }
        }
      },
      error: (httpErrorResponse: any) => {
        if (httpErrorResponse.error) {
          this.responseContent = httpErrorResponse.error;
        } else {
          this.responseContent = 'Une erreur est survenue !';
        }
        this.isUserAuthenticated = false;
        this.isSpinnerToggled = this.toggleSpinner(false, 0);
      },
      complete: () => {
      },
    });
  }

  // Logout user
  logout(): any {
    var observable = this.userService.logoutUser();

    // Send request
    observable.subscribe({
      next: () => {
        this.storageService.cleanSession();
        this.isUserAuthenticated = false;
        location.reload(); // refresh page
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

  // Create user account
  createUserAcct(): any {
    // Show spinner
    this.isSpinnerToggled = this.toggleSpinner(true, 1);

    var observable = this.userService.registerUser(this.username, this.email,
      this.password, this.confirmPass);

    // Send request
    observable.subscribe({
      next: (event: any) => {
        if (event instanceof HttpResponse) {
          this.responseContent = event.body;
          if (event.ok) {
            this.isSpinnerToggled = this.toggleSpinner(false, 1);
          }
        }
      },
      error: (httpErrorResponse: any) => {
        if (httpErrorResponse.error) {
          this.responseContent = httpErrorResponse.error;
        } else {
          this.responseContent = 'Une erreur est survenue !';
        }
        this.isSpinnerToggled = this.toggleSpinner(false, 1);
      },
      complete: () => {
      },
    });
  }

  // Toggle spinner
  private toggleSpinner(isVisible: boolean, spinnerIndx: number): boolean {
    var spinner = document.getElementsByClassName('spinner-border')[spinnerIndx] as HTMLObjectElement;
    isVisible ? spinner.style.display = 'inline-block' : spinner.style.display = 'none';

    return isVisible;
  }
}
