/*
 * wars of warp - an implementation of warpwar
 *
 * Copyright (c) 2022 Michael D Henderson
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package server

import (
	"github.com/mdhender/wow/internal/way"
	"net/http"
)

// Routes creates a new http.ServeMux and adds all the routes used by the server.
// It returns a mux which may be used directly.
func (s *Server) Routes(public string) http.Handler {
	s.router = way.NewRouter()
	s.router.Handle("GET", "/wow", s.handleIndex(public, 40, 40))
	s.router.Handle("GET", "/wow/map/color", s.handleStandardMap(public, true))
	s.router.Handle("GET", "/wow/map/mono", s.handleStandardMap(public, false))
	s.router.HandleFunc("GET", "/wow/map/random", s.handleRandomMap())
	s.router.HandleFunc("POST", "/wow/api/map-data", s.handlePostMapData())
	return s.router
}
