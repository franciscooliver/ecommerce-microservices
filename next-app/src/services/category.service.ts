import { Category } from "../models";

export class CategoryService {
 
  static create(): CategoryService {
    return new CategoryService();
  }

  async getCategories(): Promise<Category[]> {
    const response = await fetch(`${process.env.CATALOG_API_URL}/category`, {
      next: {
        revalidate: 1,
      },
    }); //revalidate on demand
    return response.json();
  }
}
