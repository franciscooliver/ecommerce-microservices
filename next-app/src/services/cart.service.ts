import { cookies } from "next/headers";
import { ReadonlyRequestCookies } from "next/dist/server/web/spec-extension/adapters/request-cookies";
import { ProductService } from "./product.service";
import { number } from "yup";

export type CartItem = {
  product_id: string;
  quantity: number;
  total: number;
};

export type Cart = {
  items: CartItem[];
  total: number;
};

export class CartService {
  cookieStore: ReadonlyRequestCookies;

  constructor(private productService: ProductService) {
    this.cookieStore = cookies();
  }

  async addToCart(input: { product_id: string; quantity: number }) {
    const cartString = this.cookieStore.get("cart")?.value;

    if (!cartString) {
      this.cookieStore.set("cart", JSON.stringify({ items: [], total: 0 }));
    }

    const cart: Cart = cartString ? JSON.parse(cartString) : { items: [], total: 0 };

    const { product_id, quantity } = input;

    const product = await this.productService.getProduct(product_id);

    const productPrice = product.price * quantity;

    cart.items.push({
      product_id,
      quantity,
      total: productPrice,
    });
    cart.total += productPrice;

    this.cookieStore.set("cart", JSON.stringify(cart));
  }

  async removeItemFromCart(index: number) {
    const cartRaw = this.cookieStore.get("cart")?.value;

    const cart: Cart = cartRaw ? JSON.parse(cartRaw) : { items: [] };

    cart.items.splice(index, 1);

    let total: number = 0;
    for(const item of cart.items) {
      const product = await this.productService.getProduct(item.product_id);
      total += item.quantity * product.price;
    }

    console.log(total);
    cart.total = total;

    this.cookieStore.set("cart", JSON.stringify(cart));
  }

  getCart() {
    const cartRaw = this.cookieStore.get("cart")?.value;

    const cart: Cart = cartRaw ? JSON.parse(cartRaw) : { items: [], total: 0 };

    return cart;
  }

  clearCart() {
    this.cookieStore.delete("cart");
  }
}

export class CartServiceFactory {
  static create() {
    const productService = new ProductService();
    return new CartService(productService);
  }
}
