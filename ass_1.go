package main

import "fmt"

func main() {
	newFilm := NewFilm(Interstellar{}, Drama{}, PopularActors{}, HighRating{})
	fmt.Println(" ")
	newFilm.Description()
	newFilm.ShowCategory()
	newFilm.ShowCaste()
	newFilm.ShowRating()
	fmt.Println(" ")
	newFilm.SetFilmName(TheMask{})
	newFilm.SetCategory(Comedy{})
	newFilm.SetFilmCast(PopularActors{})
	newFilm.SetRating(HighRating{})
	newFilm.Description()
	newFilm.ShowCategory()
	newFilm.ShowCaste()
	newFilm.ShowRating()
}

type FilmName interface {
	Description()
}
type CategoryOfFilm interface {
	ShowCategory()
}

type ActorСaste interface {
	ShowCaste()
}
type Rating interface {
	ShowRating()
}


// Category of film

type Horror struct {
}

func (hr Horror) ShowCategory() {
	fmt.Println("Film category: Horror.")
}

type Comedy struct {
}

func (cm Comedy) ShowCategory() {
	fmt.Println("Film category: Comedy.")
}

type Action struct {
}

func (ac Action) ShowCategory() {
	fmt.Println("Film category: Action.")
}

type Fantasy struct {
}

func (ft Fantasy) ShowCategory() {
	fmt.Println("Film category: Fantasy.")
}

type Drama struct {
}

func (ad Drama) ShowCategory() {
	fmt.Println("Film category: Drama.")
}

// Actor caste of film

type PopularActors struct {
}

func (pa PopularActors) ShowCaste() {
	fmt.Println("The film stars popular actors.")
}

type JuniorActors struct {
}

func (ja JuniorActors) ShowCaste() {
	fmt.Println("The film stars junior actors.")
}

type UnknownActors struct {
}

func (ua UnknownActors) ShowCaste() {
	fmt.Println("The film stars unknown actors.")
}

// Rating of film

type HighRating struct {
}

func (hr HighRating) ShowRating() {
	fmt.Println("This film has a high rating!")
}

type MediumRating struct {
}

func (mr MediumRating) ShowRating() {
	fmt.Println("This film has a medium rating!")
}

type LowRating struct {
}

func (lr LowRating) ShowRating() {
	fmt.Println("This film has a low rating!")
}

// Film name and description

type Interstellar struct {
}

func (i Interstellar) Description() {
	fmt.Println("Interstellar: " + "When drought, dust storms and plant extinctions lead humanity into a food crisis, a team of explorers and scientists travels through a wormhole (which supposedly connects regions of space-time across a great distance) to surpass previous limitations on human space travel and find a planet with conditions suitable for humanity.")
}

type TheConjuring struct {
	
}

func (tc TheConjuring) Description() {
	fmt.Println("The Cunjuring: " + "The frightening true story of demonologists Ed and Lorraine Warren, world-renowned paranormal detectives. On their secluded farm in New England, a family is unexpectedly attacked by a dark spirit. Ed and Lorraine come to the rescue as they face the most frightening case in their practice.")
}

type TheMask struct {

}

func (tm TheMask) Description() {
	fmt.Println("The Mask: " + "A modest and shy bank clerk feels insecure around pretty girls and people in general. As fate would have it, he gets a magic mask, and Stanley Ipkis gets the ability to turn into an invulnerable cartoon creature with a mischievous character.")
}

type Hachiko struct {
	
}

func (h Hachiko) Description() {
	fmt.Println("Hachico: " + "The story is based on a true story that happened in Japan and shocked the world. One day, on his way back from work, a college professor found a cute little Akita-inu puppy at a train station. Professor and Hachiko became true friends. Every day the dog accompanied and met him at the station. And even the loss of his master did not stop the dog in his hope that his friend would return.")
}

type Tenet struct {

}

func (t Tenet) Description() {
	fmt.Println("The Tenet: " + "After the terrorist attack at the Kiev Opera House, a CIA agent teams up with British intelligence to confront a Russian oligarch who made his fortune in the arms trade. To do so, the agents use time inversion, a future technology that allows time to go backwards.")
}


type Film struct {
	FilmName       FilmName
	CategoryOfFilm CategoryOfFilm
	ActorСaste     ActorСaste
	Rating         Rating

}

func NewFilm(fn FilmName, cof CategoryOfFilm, ac ActorСaste, rt Rating) *Film {
	return &Film{
		FilmName:       fn,
		CategoryOfFilm: cof,
		ActorСaste:     ac,
		Rating:         rt,
	}
}

func (f *Film) ShowCategory() {
	f.CategoryOfFilm.ShowCategory()
}

func (f *Film) SetCategory(cof_1 CategoryOfFilm) {
	f.CategoryOfFilm = cof_1
}

func (f *Film) ShowCaste() {
	f.ActorСaste.ShowCaste()
}

func (f *Film) SetFilmCast(ac_1 ActorСaste) {
	f.ActorСaste = ac_1
}

func (f *Film) ShowRating() {
	f.Rating.ShowRating()
}

func (f *Film) SetRating(rt_1 Rating) {
	f.Rating = rt_1
}
func (f* Film) Description() {
	f.FilmName.Description()
}
func (f*Film) SetFilmName(fn_1 FilmName) {
	f.FilmName = fn_1
}
