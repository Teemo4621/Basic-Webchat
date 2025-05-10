import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";

export function middleware(request: NextRequest) {
    const token = request.cookies.get("auth.token");
    if (request.nextUrl.pathname.startsWith("/login")) {
        if (token) {
            return NextResponse.redirect(new URL("/rooms", request.url));
        }
    }

    if (request.nextUrl.pathname.startsWith("/register")) {
        if (token) {
            return NextResponse.redirect(new URL("/rooms", request.url));
        }
    }

    if (request.nextUrl.pathname.startsWith("/rooms")) {
        if (!token) {
            return NextResponse.redirect(new URL("/login", request.url));
        }
    }
}

export const config = {
  matcher: ["/login", "/register", "/rooms/:path*"],
};
