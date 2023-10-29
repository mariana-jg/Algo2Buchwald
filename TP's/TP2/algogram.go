package main

import (
	errores "algogram/errores"
	TDAHash "algogram/hash"
	TDALista "algogram/lista"
	posteos "algogram/posteos"
	users "algogram/usuarios"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	LOGIN        = "login"
	LOGOUT       = "logout"
	PUBLICARPOST = "publicar"
	VERPROXIMO   = "ver_siguiente_feed"
	LIKE         = "likear_post"
	MOSTRARLIKES = "mostrar_likes"
	USERVACIO    = "No User"
	COMANDO      = 0
	PEDIDO       = 1
)

const (
	SALUDO             = "Hola"
	DESPEDIDA          = "Adios"
	CONFIRMACIONPOSTEO = "Post publicado"
	CONFIRMACIONLIKE   = "Post likeado"
)

func login(enLinea *users.Usuario, usuarios TDAHash.Diccionario[string, users.Usuario], txt string) {
	if *enLinea != usuarios.Obtener(USERVACIO) {
		fmt.Println(errores.UsuarioLoggeado{}.Error())
		return
	}
	if !usuarios.Pertenece(txt) {
		fmt.Println(errores.UsuarioInexistente{}.Error())
		return
	}
	*enLinea = usuarios.Obtener(txt)
	user := *enLinea
	fmt.Printf("%s %s\n", SALUDO, user.ObtenerNombreUsuario())
}

func logout(enLinea *users.Usuario, usuarios TDAHash.Diccionario[string, users.Usuario]) {
	if *enLinea != usuarios.Obtener(USERVACIO) {
		*enLinea = usuarios.Obtener(USERVACIO)
		fmt.Println(DESPEDIDA)
		return
	}
	fmt.Println(errores.SinUsuarioEnLinea{}.Error())
}

func publicarPosteo(enLinea *users.Usuario, usuarios TDAHash.Diccionario[string, users.Usuario],
	dicPosteo TDAHash.Diccionario[int, posteos.Post], id *int, txt string) {
	if *enLinea == usuarios.Obtener(USERVACIO) {
		fmt.Println(errores.SinUsuarioEnLinea{}.Error())
		return
	}
	user := *enLinea
	posteoNuevo := posteos.CrearPost(user.ObtenerNombreUsuario(), user.ObtenerOrdenUsuario(), *id, txt)
	dicPosteo.Guardar(*id, posteoNuevo)
	iter := usuarios.Iterador()
	for iter.HaySiguiente() {
		clave, valor := iter.VerActual()
		if clave != user.ObtenerNombreUsuario() {
			valor.GuardarEnFeed(posteoNuevo)
		}
		iter.Siguiente()
	}
	*id++
	fmt.Println(CONFIRMACIONPOSTEO)
}

func verProximo(enLinea *users.Usuario, usuarios TDAHash.Diccionario[string, users.Usuario]) {
	if *enLinea == usuarios.Obtener(USERVACIO) {
		fmt.Println(errores.ErrorSinUsuarioSinPosteo{}.Error())
		return
	}
	user := *enLinea
	proximo := user.VerProximo()
	if proximo == nil {
		fmt.Println(errores.ErrorSinUsuarioSinPosteo{}.Error())
		return
	}
	proximoPost := *proximo
	nombrepost := proximoPost.ObtenerPosteador()
	postcontenido := proximoPost.ObtenerContenidoDePost()
	idpost := proximoPost.ObtenerIDPost()
	likespost := proximoPost.ObtenerLikes()
	fmt.Printf("Post ID %d\n%s dijo: %s\nLikes: %d\n", idpost, nombrepost, postcontenido, likespost)
}

func likear(enLinea users.Usuario, comandos []string, dicPosteo TDAHash.Diccionario[int, posteos.Post],
	usuarios TDAHash.Diccionario[string, users.Usuario]) {
	if enLinea == usuarios.Obtener(USERVACIO) {
		fmt.Println(errores.ErrorSinUsuarioPostInexistente{}.Error())
		return
	}
	id, _ := strconv.Atoi(comandos[PEDIDO])
	if !dicPosteo.Pertenece(id) {
		fmt.Println(errores.ErrorSinUsuarioPostInexistente{}.Error())
		return
	}
	posteo := dicPosteo.Obtener(id)
	posteo.LikearPost(enLinea.ObtenerNombreUsuario())
	fmt.Println(CONFIRMACIONLIKE)
}

func imprimirLikes(dieronLike TDALista.Lista[string], likes int) {
	fmt.Printf("El post tiene %d likes:\n", likes)
	iter := dieronLike.Iterador()
	for iter.HaySiguiente() {
		clave := iter.VerActual()
		fmt.Printf("\t%s\n", clave)
		iter.Siguiente()
	}
}

func mostrarLikes(comandos []string, dicPosteo TDAHash.Diccionario[int, posteos.Post]) {
	id, _ := strconv.Atoi(comandos[PEDIDO])
	if !dicPosteo.Pertenece(id) {
		fmt.Println(errores.ErrorPostInexistenteSinLikes{}.Error())
		return
	}
	posteo := dicPosteo.Obtener(id)
	likes := posteo.ObtenerLikes()
	if likes == 0 {
		fmt.Println(errores.ErrorPostInexistenteSinLikes{}.Error())
		return
	}
	dieronLike := posteo.ObtenerListaDeLikes()
	imprimirLikes(dieronLike, likes)
}

func ingresarComando(comandos []string, usuarios TDAHash.Diccionario[string, users.Usuario], enLinea *users.Usuario,
	id *int, dicPosteo TDAHash.Diccionario[int, posteos.Post]) {
	txt := strings.Join(comandos[PEDIDO:], " ")
	switch comandos[COMANDO] {
	case LOGIN:
		login(enLinea, usuarios, txt)
	case LOGOUT:
		logout(enLinea, usuarios)
	case PUBLICARPOST:
		publicarPosteo(enLinea, usuarios, dicPosteo, id, txt)
	case VERPROXIMO:
		verProximo(enLinea, usuarios)
	case LIKE:
		likear(*enLinea, comandos, dicPosteo, usuarios)
	case MOSTRARLIKES:
		mostrarLikes(comandos, dicPosteo)
	}
}

func cargarUsuarios(archivoUsuario *os.File) TDAHash.Diccionario[string, users.Usuario] {
	dicUsuario := TDAHash.CrearHash[string, users.Usuario]()
	scanner := bufio.NewScanner(archivoUsuario)
	var contadorLinea int
	usuarioVacio := users.CrearUsuario[string](USERVACIO, -1)
	dicUsuario.Guardar(USERVACIO, usuarioVacio)
	for scanner.Scan() {
		linea := scanner.Text()
		usuario := users.CrearUsuario[string](linea, contadorLinea)
		dicUsuario.Guardar(linea, usuario)
		contadorLinea++
	}
	return dicUsuario
}

func main() {
	archivos := os.Args[1:]
	if len(archivos) != 1 {
		fmt.Println()
	}
	archivoUsuario, err := os.Open(archivos[0])
	if err != nil {
		fmt.Println(errores.ErrorLeerArchivo{}.Error())
		return
	}
	defer archivoUsuario.Close()

	dicPosteo := TDAHash.CrearHash[int, posteos.Post]()
	dicUsuario := cargarUsuarios(archivoUsuario)

	enLinea := dicUsuario.Obtener(USERVACIO)

	var id int
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		entrada := s.Text()
		comandos := strings.Split(entrada, " ")
		ingresarComando(comandos, dicUsuario, &enLinea, &id, dicPosteo)
	}
}
