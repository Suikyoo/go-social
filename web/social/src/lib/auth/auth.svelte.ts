import {type User} from '../types/user';

interface formController {
  type: string;
  visible: boolean;
  isVisible(s: string): boolean;
  setVisible(s: string): void;
  setHidden(): void;

}

let authForm: formController = $state({
  type: "login", 
  visible: false,
  isVisible(s: string): boolean {
    return (this.visible && this.type == s)
  },
  setVisible(s: string) {
    this.visible = true;
    this.type = s;
  },
  setHidden() {
    this.visible = false
  },
});

export {authForm}



