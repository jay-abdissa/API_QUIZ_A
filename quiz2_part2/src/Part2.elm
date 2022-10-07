module Part2 exposing (main) 
import Html exposing (..)
import Html.Attributes exposing (class, src)
import Html.Events exposing (onClick)
import Browser

type Msg =
    Like | Unlike

initialModel : { liked : Bool }
initialModel = 
    {
        liked = True
    }
    
viewstar : { liked : Bool } -> Html Msg
viewstar model =
    let
        buttonType = 
            if model.liked then "star_outline" else "star"
        msg =
            if model.liked then Unlike else Like
    in
    div [ class "header" ][
            span [ class "material-icons md-100", onClick msg ] [ text buttonType ] ]
view : { liked : Bool } -> Html Msg
view model =
    viewstar model
    
update : Msg -> { liked : Bool } -> { liked : Bool }
update msg model =
    case msg of
        Like ->
            { model | liked = True }
        Unlike ->
            { model | liked = False }

main : Program () { liked : Bool } Msg
main =
    Browser.sandbox
    {
        init = initialModel
        ,view = view
        ,update = update
    }