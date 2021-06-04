// Should have access modifiers

class Animal {
    // No accessibility modifier
    constructor(breed, name) {
      this.animalName = name;
    }
    
    // No accessibility modifier
    animalName: string; 
  
    // No accessibility modifier
    get name(): string {
      return this.animalName;
    }
  
    // No accessibility modifier
    set name(value: string) {
      this.animalName = value;
    }
    
    // method
    walk() {}
  }