package gifstatic

type emojiMeta struct {
	Keywords []string
	Name     string
	Category string
}

var emojiMetaForRunes = map[string]*emojiMeta{
	"😀": &emojiMeta{
		Keywords: []string{"face", "smile", "happy", "joy", ":D", "grin"},
		Name:     "grinning",
		Category: "people",
	},
	"😬": &emojiMeta{
		Keywords: []string{"face", "grimace", "teeth"},
		Name:     "grimacing",
		Category: "people",
	},
	"😁": &emojiMeta{
		Keywords: []string{"face", "happy", "smile", "joy", "kawaii"},
		Name:     "grin",
		Category: "people",
	},
	"😂": &emojiMeta{
		Keywords: []string{"face", "cry", "tears", "weep", "happy", "happytears", "haha"},
		Name:     "joy",
		Category: "people",
	},
	"🤣": &emojiMeta{
		Keywords: []string{"face", "rolling", "floor", "laughing", "lol", "haha"},
		Name:     "rofl",
		Category: "people",
	},
	"🥳": &emojiMeta{
		Keywords: []string{"face", "celebration", "woohoo"},
		Name:     "partying",
		Category: "people",
	},
	"😃": &emojiMeta{
		Keywords: []string{"face", "happy", "joy", "haha", ":D", ":)", "smile", "funny"},
		Name:     "smiley",
		Category: "people",
	},
	"😄": &emojiMeta{
		Keywords: []string{"face", "happy", "joy", "funny", "haha", "laugh", "like", ":D", ":)"},
		Name:     "smile",
		Category: "people",
	},
	"😅": &emojiMeta{
		Keywords: []string{"face", "hot", "happy", "laugh", "sweat", "smile", "relief"},
		Name:     "sweat_smile",
		Category: "people",
	},
	"😆": &emojiMeta{
		Keywords: []string{"happy", "joy", "lol", "satisfied", "haha", "face", "glad", "XD", "laugh"},
		Name:     "laughing",
		Category: "people",
	},
	"😇": &emojiMeta{
		Keywords: []string{"face", "angel", "heaven", "halo"},
		Name:     "innocent",
		Category: "people",
	},
	"😉": &emojiMeta{
		Keywords: []string{"face", "happy", "mischievous", "secret", ";)", "smile", "eye"},
		Name:     "wink",
		Category: "people",
	},
	"😊": &emojiMeta{
		Keywords: []string{"face", "smile", "happy", "flushed", "crush", "embarrassed", "shy", "joy"},
		Name:     "blush",
		Category: "people",
	},
	"🙂": &emojiMeta{
		Keywords: []string{"face", "smile"},
		Name:     "slightly_smiling_face",
		Category: "people",
	},
	"🙃": &emojiMeta{
		Keywords: []string{"face", "flipped", "silly", "smile"},
		Name:     "upside_down_face",
		Category: "people",
	},
	"☺️": &emojiMeta{
		Keywords: []string{"face", "blush", "massage", "happiness"},
		Name:     "relaxed",
		Category: "people",
	},
	"😋": &emojiMeta{
		Keywords: []string{"happy", "joy", "tongue", "smile", "face", "silly", "yummy", "nom", "delicious", "savouring"},
		Name:     "yum",
		Category: "people",
	},
	"😌": &emojiMeta{
		Keywords: []string{"face", "relaxed", "phew", "massage", "happiness"},
		Name:     "relieved",
		Category: "people",
	},
	"😍": &emojiMeta{
		Keywords: []string{"face", "love", "like", "affection", "valentines", "infatuation", "crush", "heart"},
		Name:     "heart_eyes",
		Category: "people",
	},
	"🥰": &emojiMeta{
		Keywords: []string{"face", "love", "like", "affection", "valentines", "infatuation", "crush", "hearts", "adore"},
		Name:     "smiling_face_with_three_hearts",
		Category: "people",
	},
	"😘": &emojiMeta{
		Keywords: []string{"face", "love", "like", "affection", "valentines", "infatuation", "kiss"},
		Name:     "kissing_heart",
		Category: "people",
	},
	"😗": &emojiMeta{
		Keywords: []string{"love", "like", "face", "3", "valentines", "infatuation", "kiss"},
		Name:     "kissing",
		Category: "people",
	},
	"😙": &emojiMeta{
		Keywords: []string{"face", "affection", "valentines", "infatuation", "kiss"},
		Name:     "kissing_smiling_eyes",
		Category: "people",
	},
	"😚": &emojiMeta{
		Keywords: []string{"face", "love", "like", "affection", "valentines", "infatuation", "kiss"},
		Name:     "kissing_closed_eyes",
		Category: "people",
	},
	"😜": &emojiMeta{
		Keywords: []string{"face", "prank", "childish", "playful", "mischievous", "smile", "wink", "tongue"},
		Name:     "stuck_out_tongue_winking_eye",
		Category: "people",
	},
	"🤪": &emojiMeta{
		Keywords: []string{"face", "goofy", "crazy"},
		Name:     "zany",
		Category: "people",
	},
	"🤨": &emojiMeta{
		Keywords: []string{"face", "distrust", "scepticism", "disapproval", "disbelief", "surprise"},
		Name:     "raised_eyebrow",
		Category: "people",
	},
	"🧐": &emojiMeta{
		Keywords: []string{"face", "stuffy", "wealthy"},
		Name:     "monocle",
		Category: "people",
	},
	"😝": &emojiMeta{
		Keywords: []string{"face", "prank", "playful", "mischievous", "smile", "tongue"},
		Name:     "stuck_out_tongue_closed_eyes",
		Category: "people",
	},
	"😛": &emojiMeta{
		Keywords: []string{"face", "prank", "childish", "playful", "mischievous", "smile", "tongue"},
		Name:     "stuck_out_tongue",
		Category: "people",
	},
	"🤑": &emojiMeta{
		Keywords: []string{"face", "rich", "dollar", "money"},
		Name:     "money_mouth_face",
		Category: "people",
	},
	"🤓": &emojiMeta{
		Keywords: []string{"face", "nerdy", "geek", "dork"},
		Name:     "nerd_face",
		Category: "people",
	},
	"😎": &emojiMeta{
		Keywords: []string{"face", "cool", "smile", "summer", "beach", "sunglass"},
		Name:     "sunglasses",
		Category: "people",
	},
	"🤩": &emojiMeta{
		Keywords: []string{"face", "smile", "starry", "eyes", "grinning"},
		Name:     "star_struck",
		Category: "people",
	},
	"🤡": &emojiMeta{
		Keywords: []string{"face"},
		Name:     "clown_face",
		Category: "people",
	},
	"🤠": &emojiMeta{
		Keywords: []string{"face", "cowgirl", "hat"},
		Name:     "cowboy_hat_face",
		Category: "people",
	},
	"🤗": &emojiMeta{
		Keywords: []string{"face", "smile", "hug"},
		Name:     "hugs",
		Category: "people",
	},
	"😏": &emojiMeta{
		Keywords: []string{"face", "smile", "mean", "prank", "smug", "sarcasm"},
		Name:     "smirk",
		Category: "people",
	},
	"😶": &emojiMeta{
		Keywords: []string{"face", "hellokitty"},
		Name:     "no_mouth",
		Category: "people",
	},
	"😐": &emojiMeta{
		Keywords: []string{"indifference", "meh", ":|", "neutral"},
		Name:     "neutral_face",
		Category: "people",
	},
	"😑": &emojiMeta{
		Keywords: []string{"face", "indifferent", "-_-", "meh", "deadpan"},
		Name:     "expressionless",
		Category: "people",
	},
	"😒": &emojiMeta{
		Keywords: []string{"indifference", "bored", "straight face", "serious", "sarcasm", "unimpressed", "skeptical", "dubious", "side_eye"},
		Name:     "unamused",
		Category: "people",
	},
	"🙄": &emojiMeta{
		Keywords: []string{"face", "eyeroll", "frustrated"},
		Name:     "roll_eyes",
		Category: "people",
	},
	"🤔": &emojiMeta{
		Keywords: []string{"face", "hmmm", "think", "consider"},
		Name:     "thinking",
		Category: "people",
	},
	"🤥": &emojiMeta{
		Keywords: []string{"face", "lie", "pinocchio"},
		Name:     "lying_face",
		Category: "people",
	},
	"🤭": &emojiMeta{
		Keywords: []string{"face", "whoops", "shock", "surprise"},
		Name:     "hand_over_mouth",
		Category: "people",
	},
	"🤫": &emojiMeta{
		Keywords: []string{"face", "quiet", "shhh"},
		Name:     "shushing",
		Category: "people",
	},
	"🤬": &emojiMeta{
		Keywords: []string{"face", "swearing", "cursing", "cussing", "profanity", "expletive"},
		Name:     "symbols_over_mouth",
		Category: "people",
	},
	"🤯": &emojiMeta{
		Keywords: []string{"face", "shocked", "mind", "blown"},
		Name:     "exploding_head",
		Category: "people",
	},
	"😳": &emojiMeta{
		Keywords: []string{"face", "blush", "shy", "flattered"},
		Name:     "flushed",
		Category: "people",
	},
	"😞": &emojiMeta{
		Keywords: []string{"face", "sad", "upset", "depressed", ":("},
		Name:     "disappointed",
		Category: "people",
	},
	"😟": &emojiMeta{
		Keywords: []string{"face", "concern", "nervous", ":("},
		Name:     "worried",
		Category: "people",
	},
	"😠": &emojiMeta{
		Keywords: []string{"mad", "face", "annoyed", "frustrated"},
		Name:     "angry",
		Category: "people",
	},
	"😡": &emojiMeta{
		Keywords: []string{"angry", "mad", "hate", "despise"},
		Name:     "rage",
		Category: "people",
	},
	"😔": &emojiMeta{
		Keywords: []string{"face", "sad", "depressed", "upset"},
		Name:     "pensive",
		Category: "people",
	},
	"😕": &emojiMeta{
		Keywords: []string{"face", "indifference", "huh", "weird", "hmmm", ":/"},
		Name:     "confused",
		Category: "people",
	},
	"🙁": &emojiMeta{
		Keywords: []string{"face", "frowning", "disappointed", "sad", "upset"},
		Name:     "slightly_frowning_face",
		Category: "people",
	},
	"☹": &emojiMeta{
		Keywords: []string{"face", "sad", "upset", "frown"},
		Name:     "frowning_face",
		Category: "people",
	},
	"😣": &emojiMeta{
		Keywords: []string{"face", "sick", "no", "upset", "oops"},
		Name:     "persevere",
		Category: "people",
	},
	"😖": &emojiMeta{
		Keywords: []string{"face", "confused", "sick", "unwell", "oops", ":S"},
		Name:     "confounded",
		Category: "people",
	},
	"😫": &emojiMeta{
		Keywords: []string{"sick", "whine", "upset", "frustrated"},
		Name:     "tired_face",
		Category: "people",
	},
	"😩": &emojiMeta{
		Keywords: []string{"face", "tired", "sleepy", "sad", "frustrated", "upset"},
		Name:     "weary",
		Category: "people",
	},
	"🥺": &emojiMeta{
		Keywords: []string{"face", "begging", "mercy"},
		Name:     "pleading",
		Category: "people",
	},
	"😤": &emojiMeta{
		Keywords: []string{"face", "gas", "phew", "proud", "pride"},
		Name:     "triumph",
		Category: "people",
	},
	"😮": &emojiMeta{
		Keywords: []string{"face", "surprise", "impressed", "wow", "whoa", ":O"},
		Name:     "open_mouth",
		Category: "people",
	},
	"😱": &emojiMeta{
		Keywords: []string{"face", "munch", "scared", "omg"},
		Name:     "scream",
		Category: "people",
	},
	"😨": &emojiMeta{
		Keywords: []string{"face", "scared", "terrified", "nervous", "oops", "huh"},
		Name:     "fearful",
		Category: "people",
	},
	"😰": &emojiMeta{
		Keywords: []string{"face", "nervous", "sweat"},
		Name:     "cold_sweat",
		Category: "people",
	},
	"😯": &emojiMeta{
		Keywords: []string{"face", "woo", "shh"},
		Name:     "hushed",
		Category: "people",
	},
	"😦": &emojiMeta{
		Keywords: []string{"face", "aw", "what"},
		Name:     "frowning",
		Category: "people",
	},
	"😧": &emojiMeta{
		Keywords: []string{"face", "stunned", "nervous"},
		Name:     "anguished",
		Category: "people",
	},
	"😢": &emojiMeta{
		Keywords: []string{"face", "tears", "sad", "depressed", "upset", ":'("},
		Name:     "cry",
		Category: "people",
	},
	"😥": &emojiMeta{
		Keywords: []string{"face", "phew", "sweat", "nervous"},
		Name:     "disappointed_relieved",
		Category: "people",
	},
	"🤤": &emojiMeta{
		Keywords: []string{"face"},
		Name:     "drooling_face",
		Category: "people",
	},
	"😪": &emojiMeta{
		Keywords: []string{"face", "tired", "rest", "nap"},
		Name:     "sleepy",
		Category: "people",
	},
	"😓": &emojiMeta{
		Keywords: []string{"face", "hot", "sad", "tired", "exercise"},
		Name:     "sweat",
		Category: "people",
	},
	"🥵": &emojiMeta{
		Keywords: []string{"face", "feverish", "heat", "red", "sweating"},
		Name:     "hot",
		Category: "people",
	},
	"🥶": &emojiMeta{
		Keywords: []string{"face", "blue", "freezing", "frozen", "frostbite", "icicles"},
		Name:     "cold",
		Category: "people",
	},
	"😭": &emojiMeta{
		Keywords: []string{"face", "cry", "tears", "sad", "upset", "depressed"},
		Name:     "sob",
		Category: "people",
	},
	"😵": &emojiMeta{
		Keywords: []string{"spent", "unconscious", "xox", "dizzy"},
		Name:     "dizzy_face",
		Category: "people",
	},
	"😲": &emojiMeta{
		Keywords: []string{"face", "xox", "surprised", "poisoned"},
		Name:     "astonished",
		Category: "people",
	},
	"🤐": &emojiMeta{
		Keywords: []string{"face", "sealed", "zipper", "secret"},
		Name:     "zipper_mouth_face",
		Category: "people",
	},
	"🤢": &emojiMeta{
		Keywords: []string{"face", "vomit", "gross", "green", "sick", "throw up", "ill"},
		Name:     "nauseated_face",
		Category: "people",
	},
	"🤧": &emojiMeta{
		Keywords: []string{"face", "gesundheit", "sneeze", "sick", "allergy"},
		Name:     "sneezing_face",
		Category: "people",
	},
	"🤮": &emojiMeta{
		Keywords: []string{"face", "sick"},
		Name:     "vomiting",
		Category: "people",
	},
	"😷": &emojiMeta{
		Keywords: []string{"face", "sick", "ill", "disease"},
		Name:     "mask",
		Category: "people",
	},
	"🤒": &emojiMeta{
		Keywords: []string{"sick", "temperature", "thermometer", "cold", "fever"},
		Name:     "face_with_thermometer",
		Category: "people",
	},
	"🤕": &emojiMeta{
		Keywords: []string{"injured", "clumsy", "bandage", "hurt"},
		Name:     "face_with_head_bandage",
		Category: "people",
	},
	"🥴": &emojiMeta{
		Keywords: []string{"face", "dizzy", "intoxicated", "tipsy", "wavy"},
		Name:     "woozy",
		Category: "people",
	},
	"😴": &emojiMeta{
		Keywords: []string{"face", "tired", "sleepy", "night", "zzz"},
		Name:     "sleeping",
		Category: "people",
	},
	"💤": &emojiMeta{
		Keywords: []string{"sleepy", "tired", "dream"},
		Name:     "zzz",
		Category: "people",
	},
	"💩": &emojiMeta{
		Keywords: []string{"hankey", "shitface", "fail", "turd", "shit"},
		Name:     "poop",
		Category: "people",
	},
	"😈": &emojiMeta{
		Keywords: []string{"devil", "horns"},
		Name:     "smiling_imp",
		Category: "people",
	},
	"👿": &emojiMeta{
		Keywords: []string{"devil", "angry", "horns"},
		Name:     "imp",
		Category: "people",
	},
	"👹": &emojiMeta{
		Keywords: []string{"monster", "red", "mask", "halloween", "scary", "creepy", "devil", "demon", "japanese", "ogre"},
		Name:     "japanese_ogre",
		Category: "people",
	},
	"👺": &emojiMeta{
		Keywords: []string{"red", "evil", "mask", "monster", "scary", "creepy", "japanese", "goblin"},
		Name:     "japanese_goblin",
		Category: "people",
	},
	"💀": &emojiMeta{
		Keywords: []string{"dead", "skeleton", "creepy", "death"},
		Name:     "skull",
		Category: "people",
	},
	"👻": &emojiMeta{
		Keywords: []string{"halloween", "spooky", "scary"},
		Name:     "ghost",
		Category: "people",
	},
	"👽": &emojiMeta{
		Keywords: []string{"UFO", "paul", "weird", "outer_space"},
		Name:     "alien",
		Category: "people",
	},
	"🤖": &emojiMeta{
		Keywords: []string{"computer", "machine", "bot"},
		Name:     "robot",
		Category: "people",
	},
	"😺": &emojiMeta{
		Keywords: []string{"animal", "cats", "happy", "smile"},
		Name:     "smiley_cat",
		Category: "people",
	},
	"😸": &emojiMeta{
		Keywords: []string{"animal", "cats", "smile"},
		Name:     "smile_cat",
		Category: "people",
	},
	"😹": &emojiMeta{
		Keywords: []string{"animal", "cats", "haha", "happy", "tears"},
		Name:     "joy_cat",
		Category: "people",
	},
	"😻": &emojiMeta{
		Keywords: []string{"animal", "love", "like", "affection", "cats", "valentines", "heart"},
		Name:     "heart_eyes_cat",
		Category: "people",
	},
	"😼": &emojiMeta{
		Keywords: []string{"animal", "cats", "smirk"},
		Name:     "smirk_cat",
		Category: "people",
	},
	"😽": &emojiMeta{
		Keywords: []string{"animal", "cats", "kiss"},
		Name:     "kissing_cat",
		Category: "people",
	},
	"🙀": &emojiMeta{
		Keywords: []string{"animal", "cats", "munch", "scared", "scream"},
		Name:     "scream_cat",
		Category: "people",
	},
	"😿": &emojiMeta{
		Keywords: []string{"animal", "tears", "weep", "sad", "cats", "upset", "cry"},
		Name:     "crying_cat_face",
		Category: "people",
	},
	"😾": &emojiMeta{
		Keywords: []string{"animal", "cats"},
		Name:     "pouting_cat",
		Category: "people",
	},
	"🤲": &emojiMeta{
		Keywords: []string{"hands", "gesture", "cupped", "prayer"},
		Name:     "palms_up",
		Category: "people",
	},
	"🙌": &emojiMeta{
		Keywords: []string{"gesture", "hooray", "yea", "celebration", "hands"},
		Name:     "raised_hands",
		Category: "people",
	},
	"👏": &emojiMeta{
		Keywords: []string{"hands", "praise", "applause", "congrats", "yay"},
		Name:     "clap",
		Category: "people",
	},
	"👋": &emojiMeta{
		Keywords: []string{"hands", "gesture", "goodbye", "solong", "farewell", "hello", "hi", "palm"},
		Name:     "wave",
		Category: "people",
	},
	"🤙": &emojiMeta{
		Keywords: []string{"hands", "gesture"},
		Name:     "call_me_hand",
		Category: "people",
	},
	"👍": &emojiMeta{
		Keywords: []string{"thumbsup", "yes", "awesome", "good", "agree", "accept", "cool", "hand", "like"},
		Name:     "+1",
		Category: "people",
	},
	"👎": &emojiMeta{
		Keywords: []string{"thumbsdown", "no", "dislike", "hand"},
		Name:     "-1",
		Category: "people",
	},
	"👊": &emojiMeta{
		Keywords: []string{"angry", "violence", "fist", "hit", "attack", "hand"},
		Name:     "facepunch",
		Category: "people",
	},
	"✊": &emojiMeta{
		Keywords: []string{"fingers", "hand", "grasp"},
		Name:     "fist",
		Category: "people",
	},
	"🤛": &emojiMeta{
		Keywords: []string{"hand", "fistbump"},
		Name:     "fist_left",
		Category: "people",
	},
	"🤜": &emojiMeta{
		Keywords: []string{"hand", "fistbump"},
		Name:     "fist_right",
		Category: "people",
	},
	"✌": &emojiMeta{
		Keywords: []string{"fingers", "ohyeah", "hand", "peace", "victory", "two"},
		Name:     "v",
		Category: "people",
	},
	"👌": &emojiMeta{
		Keywords: []string{"fingers", "limbs", "perfect", "ok", "okay"},
		Name:     "ok_hand",
		Category: "people",
	},
	"✋": &emojiMeta{
		Keywords: []string{"fingers", "stop", "highfive", "palm", "ban"},
		Name:     "raised_hand",
		Category: "people",
	},
	"🤚": &emojiMeta{
		Keywords: []string{"fingers", "raised", "backhand"},
		Name:     "raised_back_of_hand",
		Category: "people",
	},
	"👐": &emojiMeta{
		Keywords: []string{"fingers", "butterfly", "hands", "open"},
		Name:     "open_hands",
		Category: "people",
	},
	"💪": &emojiMeta{
		Keywords: []string{"arm", "flex", "hand", "summer", "strong", "biceps"},
		Name:     "muscle",
		Category: "people",
	},
	"🙏": &emojiMeta{
		Keywords: []string{"please", "hope", "wish", "namaste", "highfive"},
		Name:     "pray",
		Category: "people",
	},
	"🦶": &emojiMeta{
		Keywords: []string{"kick", "stomp"},
		Name:     "foot",
		Category: "people",
	},
	"🦵": &emojiMeta{
		Keywords: []string{"kick", "limb"},
		Name:     "leg",
		Category: "people",
	},
	"🤝": &emojiMeta{
		Keywords: []string{"agreement", "shake"},
		Name:     "handshake",
		Category: "people",
	},
	"☝": &emojiMeta{
		Keywords: []string{"hand", "fingers", "direction", "up"},
		Name:     "point_up",
		Category: "people",
	},
	"👆": &emojiMeta{
		Keywords: []string{"fingers", "hand", "direction", "up"},
		Name:     "point_up_2",
		Category: "people",
	},
	"👇": &emojiMeta{
		Keywords: []string{"fingers", "hand", "direction", "down"},
		Name:     "point_down",
		Category: "people",
	},
	"👈": &emojiMeta{
		Keywords: []string{"direction", "fingers", "hand", "left"},
		Name:     "point_left",
		Category: "people",
	},
	"👉": &emojiMeta{
		Keywords: []string{"fingers", "hand", "direction", "right"},
		Name:     "point_right",
		Category: "people",
	},
	"🖕": &emojiMeta{
		Keywords: []string{"hand", "fingers", "rude", "middle", "flipping"},
		Name:     "fu",
		Category: "people",
	},
	"🖐": &emojiMeta{
		Keywords: []string{"hand", "fingers", "palm"},
		Name:     "raised_hand_with_fingers_splayed",
		Category: "people",
	},
	"🤟": &emojiMeta{
		Keywords: []string{"hand", "fingers", "gesture"},
		Name:     "love_you",
		Category: "people",
	},
	"🤘": &emojiMeta{
		Keywords: []string{"hand", "fingers", "evil_eye", "sign_of_horns", "rock_on"},
		Name:     "metal",
		Category: "people",
	},
	"🤞": &emojiMeta{
		Keywords: []string{"good", "lucky"},
		Name:     "crossed_fingers",
		Category: "people",
	},
	"🖖": &emojiMeta{
		Keywords: []string{"hand", "fingers", "spock", "star trek"},
		Name:     "vulcan_salute",
		Category: "people",
	},
	"✍": &emojiMeta{
		Keywords: []string{"lower_left_ballpoint_pen", "stationery", "write", "compose"},
		Name:     "writing_hand",
		Category: "people",
	},
	"🤳": &emojiMeta{
		Keywords: []string{"camera", "phone"},
		Name:     "selfie",
		Category: "people",
	},
	"💅": &emojiMeta{
		Keywords: []string{"beauty", "manicure", "finger", "fashion", "nail"},
		Name:     "nail_care",
		Category: "people",
	},
	"👄": &emojiMeta{
		Keywords: []string{"mouth", "kiss"},
		Name:     "lips",
		Category: "people",
	},
	"🦷": &emojiMeta{
		Keywords: []string{"teeth", "dentist"},
		Name:     "tooth",
		Category: "people",
	},
	"👅": &emojiMeta{
		Keywords: []string{"mouth", "playful"},
		Name:     "tongue",
		Category: "people",
	},
	"👂": &emojiMeta{
		Keywords: []string{"face", "hear", "sound", "listen"},
		Name:     "ear",
		Category: "people",
	},
	"👃": &emojiMeta{
		Keywords: []string{"smell", "sniff"},
		Name:     "nose",
		Category: "people",
	},
	"👁": &emojiMeta{
		Keywords: []string{"face", "look", "see", "watch", "stare"},
		Name:     "eye",
		Category: "people",
	},
	"👀": &emojiMeta{
		Keywords: []string{"look", "watch", "stalk", "peek", "see"},
		Name:     "eyes",
		Category: "people",
	},
	"🧠": &emojiMeta{
		Keywords: []string{"smart", "intelligent"},
		Name:     "brain",
		Category: "people",
	},
	"👤": &emojiMeta{
		Keywords: []string{"user", "person", "human"},
		Name:     "bust_in_silhouette",
		Category: "people",
	},
	"👥": &emojiMeta{
		Keywords: []string{"user", "person", "human", "group", "team"},
		Name:     "busts_in_silhouette",
		Category: "people",
	},
	"🗣": &emojiMeta{
		Keywords: []string{"user", "person", "human", "sing", "say", "talk"},
		Name:     "speaking_head",
		Category: "people",
	},
	"👶": &emojiMeta{
		Keywords: []string{"child", "boy", "girl", "toddler"},
		Name:     "baby",
		Category: "people",
	},
	"🧒": &emojiMeta{
		Keywords: []string{"gender-neutral", "young"},
		Name:     "child",
		Category: "people",
	},
	"👦": &emojiMeta{
		Keywords: []string{"man", "male", "guy", "teenager"},
		Name:     "boy",
		Category: "people",
	},
	"👧": &emojiMeta{
		Keywords: []string{"female", "woman", "teenager"},
		Name:     "girl",
		Category: "people",
	},
	"🧑": &emojiMeta{
		Keywords: []string{"gender-neutral", "person"},
		Name:     "adult",
		Category: "people",
	},
	"👨": &emojiMeta{
		Keywords: []string{"mustache", "father", "dad", "guy", "classy", "sir", "moustache"},
		Name:     "man",
		Category: "people",
	},
	"👩": &emojiMeta{
		Keywords: []string{"female", "girls", "lady"},
		Name:     "woman",
		Category: "people",
	},
	"👱‍♀️": &emojiMeta{
		Keywords: []string{"woman", "female", "girl", "blonde", "person"},
		Name:     "blonde_woman",
		Category: "people",
	},
	"👱": &emojiMeta{
		Keywords: []string{"man", "male", "boy", "blonde", "guy", "person"},
		Name:     "blonde_man",
		Category: "people",
	},
	"🧔": &emojiMeta{
		Keywords: []string{"person", "bewhiskered"},
		Name:     "bearded_person",
		Category: "people",
	},
	"🧓": &emojiMeta{
		Keywords: []string{"human", "elder", "senior", "gender-neutral"},
		Name:     "older_adult",
		Category: "people",
	},
	"👴": &emojiMeta{
		Keywords: []string{"human", "male", "men", "old", "elder", "senior"},
		Name:     "older_man",
		Category: "people",
	},
	"👵": &emojiMeta{
		Keywords: []string{"human", "female", "women", "lady", "old", "elder", "senior"},
		Name:     "older_woman",
		Category: "people",
	},
	"👲": &emojiMeta{
		Keywords: []string{"male", "boy", "chinese"},
		Name:     "man_with_gua_pi_mao",
		Category: "people",
	},
	"🧕": &emojiMeta{
		Keywords: []string{"female", "hijab", "mantilla", "tichel"},
		Name:     "woman_with_headscarf",
		Category: "people",
	},
	"👳‍♀️": &emojiMeta{
		Keywords: []string{"female", "indian", "hinduism", "arabs", "woman"},
		Name:     "woman_with_turban",
		Category: "people",
	},
	"👳": &emojiMeta{
		Keywords: []string{"male", "indian", "hinduism", "arabs"},
		Name:     "man_with_turban",
		Category: "people",
	},
	"👮‍♀️": &emojiMeta{
		Keywords: []string{"woman", "police", "law", "legal", "enforcement", "arrest", "911", "female"},
		Name:     "policewoman",
		Category: "people",
	},
	"👮": &emojiMeta{
		Keywords: []string{"man", "police", "law", "legal", "enforcement", "arrest", "911"},
		Name:     "policeman",
		Category: "people",
	},
	"👷‍♀️": &emojiMeta{
		Keywords: []string{"female", "human", "wip", "build", "construction", "worker", "labor", "woman"},
		Name:     "construction_worker_woman",
		Category: "people",
	},
	"👷": &emojiMeta{
		Keywords: []string{"male", "human", "wip", "guy", "build", "construction", "worker", "labor"},
		Name:     "construction_worker_man",
		Category: "people",
	},
	"💂‍♀️": &emojiMeta{
		Keywords: []string{"uk", "gb", "british", "female", "royal", "woman"},
		Name:     "guardswoman",
		Category: "people",
	},
	"💂": &emojiMeta{
		Keywords: []string{"uk", "gb", "british", "male", "guy", "royal"},
		Name:     "guardsman",
		Category: "people",
	},
	"🕵️‍♀️": &emojiMeta{
		Keywords: []string{"human", "spy", "detective", "female", "woman"},
		Name:     "female_detective",
		Category: "people",
	},
	"🕵": &emojiMeta{
		Keywords: []string{"human", "spy", "detective"},
		Name:     "male_detective",
		Category: "people",
	},
	"👩‍⚕️": &emojiMeta{
		Keywords: []string{"doctor", "nurse", "therapist", "healthcare", "woman", "human"},
		Name:     "woman_health_worker",
		Category: "people",
	},
	"👨‍⚕️": &emojiMeta{
		Keywords: []string{"doctor", "nurse", "therapist", "healthcare", "man", "human"},
		Name:     "man_health_worker",
		Category: "people",
	},
	"👩‍🌾": &emojiMeta{
		Keywords: []string{"rancher", "gardener", "woman", "human"},
		Name:     "woman_farmer",
		Category: "people",
	},
	"👨‍🌾": &emojiMeta{
		Keywords: []string{"rancher", "gardener", "man", "human"},
		Name:     "man_farmer",
		Category: "people",
	},
	"👩‍🍳": &emojiMeta{
		Keywords: []string{"chef", "woman", "human"},
		Name:     "woman_cook",
		Category: "people",
	},
	"👨‍🍳": &emojiMeta{
		Keywords: []string{"chef", "man", "human"},
		Name:     "man_cook",
		Category: "people",
	},
	"👩‍🎓": &emojiMeta{
		Keywords: []string{"graduate", "woman", "human"},
		Name:     "woman_student",
		Category: "people",
	},
	"👨‍🎓": &emojiMeta{
		Keywords: []string{"graduate", "man", "human"},
		Name:     "man_student",
		Category: "people",
	},
	"👩‍🎤": &emojiMeta{
		Keywords: []string{"rockstar", "entertainer", "woman", "human"},
		Name:     "woman_singer",
		Category: "people",
	},
	"👨‍🎤": &emojiMeta{
		Keywords: []string{"rockstar", "entertainer", "man", "human"},
		Name:     "man_singer",
		Category: "people",
	},
	"👩‍🏫": &emojiMeta{
		Keywords: []string{"instructor", "professor", "woman", "human"},
		Name:     "woman_teacher",
		Category: "people",
	},
	"👨‍🏫": &emojiMeta{
		Keywords: []string{"instructor", "professor", "man", "human"},
		Name:     "man_teacher",
		Category: "people",
	},
	"👩‍🏭": &emojiMeta{
		Keywords: []string{"assembly", "industrial", "woman", "human"},
		Name:     "woman_factory_worker",
		Category: "people",
	},
	"👨‍🏭": &emojiMeta{
		Keywords: []string{"assembly", "industrial", "man", "human"},
		Name:     "man_factory_worker",
		Category: "people",
	},
	"👩‍💻": &emojiMeta{
		Keywords: []string{"coder", "developer", "engineer", "programmer", "software", "woman", "human", "laptop", "computer"},
		Name:     "woman_technologist",
		Category: "people",
	},
	"👨‍💻": &emojiMeta{
		Keywords: []string{"coder", "developer", "engineer", "programmer", "software", "man", "human", "laptop", "computer"},
		Name:     "man_technologist",
		Category: "people",
	},
	"👩‍💼": &emojiMeta{
		Keywords: []string{"business", "manager", "woman", "human"},
		Name:     "woman_office_worker",
		Category: "people",
	},
	"👨‍💼": &emojiMeta{
		Keywords: []string{"business", "manager", "man", "human"},
		Name:     "man_office_worker",
		Category: "people",
	},
	"👩‍🔧": &emojiMeta{
		Keywords: []string{"plumber", "woman", "human", "wrench"},
		Name:     "woman_mechanic",
		Category: "people",
	},
	"👨‍🔧": &emojiMeta{
		Keywords: []string{"plumber", "man", "human", "wrench"},
		Name:     "man_mechanic",
		Category: "people",
	},
	"👩‍🔬": &emojiMeta{
		Keywords: []string{"biologist", "chemist", "engineer", "physicist", "woman", "human"},
		Name:     "woman_scientist",
		Category: "people",
	},
	"👨‍🔬": &emojiMeta{
		Keywords: []string{"biologist", "chemist", "engineer", "physicist", "man", "human"},
		Name:     "man_scientist",
		Category: "people",
	},
	"👩‍🎨": &emojiMeta{
		Keywords: []string{"painter", "woman", "human"},
		Name:     "woman_artist",
		Category: "people",
	},
	"👨‍🎨": &emojiMeta{
		Keywords: []string{"painter", "man", "human"},
		Name:     "man_artist",
		Category: "people",
	},
	"👩‍🚒": &emojiMeta{
		Keywords: []string{"fireman", "woman", "human"},
		Name:     "woman_firefighter",
		Category: "people",
	},
	"👨‍🚒": &emojiMeta{
		Keywords: []string{"fireman", "man", "human"},
		Name:     "man_firefighter",
		Category: "people",
	},
	"👩‍✈️": &emojiMeta{
		Keywords: []string{"aviator", "plane", "woman", "human"},
		Name:     "woman_pilot",
		Category: "people",
	},
	"👨‍✈️": &emojiMeta{
		Keywords: []string{"aviator", "plane", "man", "human"},
		Name:     "man_pilot",
		Category: "people",
	},
	"👩‍🚀": &emojiMeta{
		Keywords: []string{"space", "rocket", "woman", "human"},
		Name:     "woman_astronaut",
		Category: "people",
	},
	"👨‍🚀": &emojiMeta{
		Keywords: []string{"space", "rocket", "man", "human"},
		Name:     "man_astronaut",
		Category: "people",
	},
	"👩‍⚖️": &emojiMeta{
		Keywords: []string{"justice", "court", "woman", "human"},
		Name:     "woman_judge",
		Category: "people",
	},
	"👨‍⚖️": &emojiMeta{
		Keywords: []string{"justice", "court", "man", "human"},
		Name:     "man_judge",
		Category: "people",
	},
	"🦸‍♀️": &emojiMeta{
		Keywords: []string{"woman", "female", "good", "heroine", "superpowers"},
		Name:     "woman_superhero",
		Category: "people",
	},
	"🦸‍♂️": &emojiMeta{
		Keywords: []string{"man", "male", "good", "hero", "superpowers"},
		Name:     "man_superhero",
		Category: "people",
	},
	"🦹‍♀️": &emojiMeta{
		Keywords: []string{"woman", "female", "evil", "bad", "criminal", "heroine", "superpowers"},
		Name:     "woman_supervillain",
		Category: "people",
	},
	"🦹‍♂️": &emojiMeta{
		Keywords: []string{"man", "male", "evil", "bad", "criminal", "hero", "superpowers"},
		Name:     "man_supervillain",
		Category: "people",
	},
	"🤶": &emojiMeta{
		Keywords: []string{"woman", "female", "xmas", "mother christmas"},
		Name:     "mrs_claus",
		Category: "people",
	},
	"🎅": &emojiMeta{
		Keywords: []string{"festival", "man", "male", "xmas", "father christmas"},
		Name:     "santa",
		Category: "people",
	},
	"🧙‍♀️": &emojiMeta{
		Keywords: []string{"woman", "female", "mage", "witch"},
		Name:     "sorceress",
		Category: "people",
	},
	"🧙‍♂️": &emojiMeta{
		Keywords: []string{"man", "male", "mage", "sorcerer"},
		Name:     "wizard",
		Category: "people",
	},
	"🧝‍♀️": &emojiMeta{
		Keywords: []string{"woman", "female"},
		Name:     "woman_elf",
		Category: "people",
	},
	"🧝‍♂️": &emojiMeta{
		Keywords: []string{"man", "male"},
		Name:     "man_elf",
		Category: "people",
	},
	"🧛‍♀️": &emojiMeta{
		Keywords: []string{"woman", "female"},
		Name:     "woman_vampire",
		Category: "people",
	},
	"🧛‍♂️": &emojiMeta{
		Keywords: []string{"man", "male", "dracula"},
		Name:     "man_vampire",
		Category: "people",
	},
	"🧟‍♀️": &emojiMeta{
		Keywords: []string{"woman", "female", "undead", "walking dead"},
		Name:     "woman_zombie",
		Category: "people",
	},
	"🧟‍♂️": &emojiMeta{
		Keywords: []string{"man", "male", "dracula", "undead", "walking dead"},
		Name:     "man_zombie",
		Category: "people",
	},
	"🧞‍♀️": &emojiMeta{
		Keywords: []string{"woman", "female"},
		Name:     "woman_genie",
		Category: "people",
	},
	"🧞‍♂️": &emojiMeta{
		Keywords: []string{"man", "male"},
		Name:     "man_genie",
		Category: "people",
	},
	"🧜‍♀️": &emojiMeta{
		Keywords: []string{"woman", "female", "merwoman", "ariel"},
		Name:     "mermaid",
		Category: "people",
	},
	"🧜‍♂️": &emojiMeta{
		Keywords: []string{"man", "male", "triton"},
		Name:     "merman",
		Category: "people",
	},
	"🧚‍♀️": &emojiMeta{
		Keywords: []string{"woman", "female"},
		Name:     "woman_fairy",
		Category: "people",
	},
	"🧚‍♂️": &emojiMeta{
		Keywords: []string{"man", "male"},
		Name:     "man_fairy",
		Category: "people",
	},
	"👼": &emojiMeta{
		Keywords: []string{"heaven", "wings", "halo"},
		Name:     "angel",
		Category: "people",
	},
	"🤰": &emojiMeta{
		Keywords: []string{"baby"},
		Name:     "pregnant_woman",
		Category: "people",
	},
	"🤱": &emojiMeta{
		Keywords: []string{"nursing", "baby"},
		Name:     "breastfeeding",
		Category: "people",
	},
	"👸": &emojiMeta{
		Keywords: []string{"girl", "woman", "female", "blond", "crown", "royal", "queen"},
		Name:     "princess",
		Category: "people",
	},
	"🤴": &emojiMeta{
		Keywords: []string{"boy", "man", "male", "crown", "royal", "king"},
		Name:     "prince",
		Category: "people",
	},
	"👰": &emojiMeta{
		Keywords: []string{"couple", "marriage", "wedding", "woman", "bride"},
		Name:     "bride_with_veil",
		Category: "people",
	},
	"🤵": &emojiMeta{
		Keywords: []string{"couple", "marriage", "wedding", "groom"},
		Name:     "man_in_tuxedo",
		Category: "people",
	},
	"🏃‍♀️": &emojiMeta{
		Keywords: []string{"woman", "walking", "exercise", "race", "running", "female"},
		Name:     "running_woman",
		Category: "people",
	},
	"🏃": &emojiMeta{
		Keywords: []string{"man", "walking", "exercise", "race", "running"},
		Name:     "running_man",
		Category: "people",
	},
	"🚶‍♀️": &emojiMeta{
		Keywords: []string{"human", "feet", "steps", "woman", "female"},
		Name:     "walking_woman",
		Category: "people",
	},
	"🚶": &emojiMeta{
		Keywords: []string{"human", "feet", "steps"},
		Name:     "walking_man",
		Category: "people",
	},
	"💃": &emojiMeta{
		Keywords: []string{"female", "girl", "woman", "fun"},
		Name:     "dancer",
		Category: "people",
	},
	"🕺": &emojiMeta{
		Keywords: []string{"male", "boy", "fun", "dancer"},
		Name:     "man_dancing",
		Category: "people",
	},
	"👯": &emojiMeta{
		Keywords: []string{"female", "bunny", "women", "girls"},
		Name:     "dancing_women",
		Category: "people",
	},
	"👯‍♂️": &emojiMeta{
		Keywords: []string{"male", "bunny", "men", "boys"},
		Name:     "dancing_men",
		Category: "people",
	},
	"👫": &emojiMeta{
		Keywords: []string{"pair", "people", "human", "love", "date", "dating", "like", "affection", "valentines", "marriage"},
		Name:     "couple",
		Category: "people",
	},
	"👬": &emojiMeta{
		Keywords: []string{"pair", "couple", "love", "like", "bromance", "friendship", "people", "human"},
		Name:     "two_men_holding_hands",
		Category: "people",
	},
	"👭": &emojiMeta{
		Keywords: []string{"pair", "friendship", "couple", "love", "like", "female", "people", "human"},
		Name:     "two_women_holding_hands",
		Category: "people",
	},
	"🙇‍♀️": &emojiMeta{
		Keywords: []string{"woman", "female", "girl"},
		Name:     "bowing_woman",
		Category: "people",
	},
	"🙇": &emojiMeta{
		Keywords: []string{"man", "male", "boy"},
		Name:     "bowing_man",
		Category: "people",
	},
	"🤦‍♂️": &emojiMeta{
		Keywords: []string{"man", "male", "boy", "disbelief"},
		Name:     "man_facepalming",
		Category: "people",
	},
	"🤦‍♀️": &emojiMeta{
		Keywords: []string{"woman", "female", "girl", "disbelief"},
		Name:     "woman_facepalming",
		Category: "people",
	},
	"🤷": &emojiMeta{
		Keywords: []string{"woman", "female", "girl", "confused", "indifferent", "doubt"},
		Name:     "woman_shrugging",
		Category: "people",
	},
	"🤷‍♂️": &emojiMeta{
		Keywords: []string{"man", "male", "boy", "confused", "indifferent", "doubt"},
		Name:     "man_shrugging",
		Category: "people",
	},
	"💁": &emojiMeta{
		Keywords: []string{"female", "girl", "woman", "human", "information"},
		Name:     "tipping_hand_woman",
		Category: "people",
	},
	"💁‍♂️": &emojiMeta{
		Keywords: []string{"male", "boy", "man", "human", "information"},
		Name:     "tipping_hand_man",
		Category: "people",
	},
	"🙅": &emojiMeta{
		Keywords: []string{"female", "girl", "woman", "nope"},
		Name:     "no_good_woman",
		Category: "people",
	},
	"🙅‍♂️": &emojiMeta{
		Keywords: []string{"male", "boy", "man", "nope"},
		Name:     "no_good_man",
		Category: "people",
	},
	"🙆": &emojiMeta{
		Keywords: []string{"women", "girl", "female", "pink", "human", "woman"},
		Name:     "ok_woman",
		Category: "people",
	},
	"🙆‍♂️": &emojiMeta{
		Keywords: []string{"men", "boy", "male", "blue", "human", "man"},
		Name:     "ok_man",
		Category: "people",
	},
	"🙋": &emojiMeta{
		Keywords: []string{"female", "girl", "woman"},
		Name:     "raising_hand_woman",
		Category: "people",
	},
	"🙋‍♂️": &emojiMeta{
		Keywords: []string{"male", "boy", "man"},
		Name:     "raising_hand_man",
		Category: "people",
	},
	"🙎": &emojiMeta{
		Keywords: []string{"female", "girl", "woman"},
		Name:     "pouting_woman",
		Category: "people",
	},
	"🙎‍♂️": &emojiMeta{
		Keywords: []string{"male", "boy", "man"},
		Name:     "pouting_man",
		Category: "people",
	},
	"🙍": &emojiMeta{
		Keywords: []string{"female", "girl", "woman", "sad", "depressed", "discouraged", "unhappy"},
		Name:     "frowning_woman",
		Category: "people",
	},
	"🙍‍♂️": &emojiMeta{
		Keywords: []string{"male", "boy", "man", "sad", "depressed", "discouraged", "unhappy"},
		Name:     "frowning_man",
		Category: "people",
	},
	"💇": &emojiMeta{
		Keywords: []string{"female", "girl", "woman"},
		Name:     "haircut_woman",
		Category: "people",
	},
	"💇‍♂️": &emojiMeta{
		Keywords: []string{"male", "boy", "man"},
		Name:     "haircut_man",
		Category: "people",
	},
	"💆": &emojiMeta{
		Keywords: []string{"female", "girl", "woman", "head"},
		Name:     "massage_woman",
		Category: "people",
	},
	"💆‍♂️": &emojiMeta{
		Keywords: []string{"male", "boy", "man", "head"},
		Name:     "massage_man",
		Category: "people",
	},
	"🧖‍♀️": &emojiMeta{
		Keywords: []string{"female", "woman", "spa", "steamroom", "sauna"},
		Name:     "woman_in_steamy_room",
		Category: "people",
	},
	"🧖‍♂️": &emojiMeta{
		Keywords: []string{"male", "man", "spa", "steamroom", "sauna"},
		Name:     "man_in_steamy_room",
		Category: "people",
	},
	"💑": &emojiMeta{
		Keywords: []string{"pair", "love", "like", "affection", "human", "dating", "valentines", "marriage"},
		Name:     "couple_with_heart_woman_man",
		Category: "people",
	},
	"👩‍❤️‍👩": &emojiMeta{
		Keywords: []string{"pair", "love", "like", "affection", "human", "dating", "valentines", "marriage"},
		Name:     "couple_with_heart_woman_woman",
		Category: "people",
	},
	"👨‍❤️‍👨": &emojiMeta{
		Keywords: []string{"pair", "love", "like", "affection", "human", "dating", "valentines", "marriage"},
		Name:     "couple_with_heart_man_man",
		Category: "people",
	},
	"💏": &emojiMeta{
		Keywords: []string{"pair", "valentines", "love", "like", "dating", "marriage"},
		Name:     "couplekiss_man_woman",
		Category: "people",
	},
	"👩‍❤️‍💋‍👩": &emojiMeta{
		Keywords: []string{"pair", "valentines", "love", "like", "dating", "marriage"},
		Name:     "couplekiss_woman_woman",
		Category: "people",
	},
	"👨‍❤️‍💋‍👨": &emojiMeta{
		Keywords: []string{"pair", "valentines", "love", "like", "dating", "marriage"},
		Name:     "couplekiss_man_man",
		Category: "people",
	},
	"👪": &emojiMeta{
		Keywords: []string{"home", "parents", "child", "mom", "dad", "father", "mother", "people", "human"},
		Name:     "family_man_woman_boy",
		Category: "people",
	},
	"👨‍👩‍👧": &emojiMeta{
		Keywords: []string{"home", "parents", "people", "human", "child"},
		Name:     "family_man_woman_girl",
		Category: "people",
	},
	"👨‍👩‍👧‍👦": &emojiMeta{
		Keywords: []string{"home", "parents", "people", "human", "children"},
		Name:     "family_man_woman_girl_boy",
		Category: "people",
	},
	"👨‍👩‍👦‍👦": &emojiMeta{
		Keywords: []string{"home", "parents", "people", "human", "children"},
		Name:     "family_man_woman_boy_boy",
		Category: "people",
	},
	"👨‍👩‍👧‍👧": &emojiMeta{
		Keywords: []string{"home", "parents", "people", "human", "children"},
		Name:     "family_man_woman_girl_girl",
		Category: "people",
	},
	"👩‍👩‍👦": &emojiMeta{
		Keywords: []string{"home", "parents", "people", "human", "children"},
		Name:     "family_woman_woman_boy",
		Category: "people",
	},
	"👩‍👩‍👧": &emojiMeta{
		Keywords: []string{"home", "parents", "people", "human", "children"},
		Name:     "family_woman_woman_girl",
		Category: "people",
	},
	"👩‍👩‍👧‍👦": &emojiMeta{
		Keywords: []string{"home", "parents", "people", "human", "children"},
		Name:     "family_woman_woman_girl_boy",
		Category: "people",
	},
	"👩‍👩‍👦‍👦": &emojiMeta{
		Keywords: []string{"home", "parents", "people", "human", "children"},
		Name:     "family_woman_woman_boy_boy",
		Category: "people",
	},
	"👩‍👩‍👧‍👧": &emojiMeta{
		Keywords: []string{"home", "parents", "people", "human", "children"},
		Name:     "family_woman_woman_girl_girl",
		Category: "people",
	},
	"👨‍👨‍👦": &emojiMeta{
		Keywords: []string{"home", "parents", "people", "human", "children"},
		Name:     "family_man_man_boy",
		Category: "people",
	},
	"👨‍👨‍👧": &emojiMeta{
		Keywords: []string{"home", "parents", "people", "human", "children"},
		Name:     "family_man_man_girl",
		Category: "people",
	},
	"👨‍👨‍👧‍👦": &emojiMeta{
		Keywords: []string{"home", "parents", "people", "human", "children"},
		Name:     "family_man_man_girl_boy",
		Category: "people",
	},
	"👨‍👨‍👦‍👦": &emojiMeta{
		Keywords: []string{"home", "parents", "people", "human", "children"},
		Name:     "family_man_man_boy_boy",
		Category: "people",
	},
	"👨‍👨‍👧‍👧": &emojiMeta{
		Keywords: []string{"home", "parents", "people", "human", "children"},
		Name:     "family_man_man_girl_girl",
		Category: "people",
	},
	"👩‍👦": &emojiMeta{
		Keywords: []string{"home", "parent", "people", "human", "child"},
		Name:     "family_woman_boy",
		Category: "people",
	},
	"👩‍👧": &emojiMeta{
		Keywords: []string{"home", "parent", "people", "human", "child"},
		Name:     "family_woman_girl",
		Category: "people",
	},
	"👩‍👧‍👦": &emojiMeta{
		Keywords: []string{"home", "parent", "people", "human", "children"},
		Name:     "family_woman_girl_boy",
		Category: "people",
	},
	"👩‍👦‍👦": &emojiMeta{
		Keywords: []string{"home", "parent", "people", "human", "children"},
		Name:     "family_woman_boy_boy",
		Category: "people",
	},
	"👩‍👧‍👧": &emojiMeta{
		Keywords: []string{"home", "parent", "people", "human", "children"},
		Name:     "family_woman_girl_girl",
		Category: "people",
	},
	"👨‍👦": &emojiMeta{
		Keywords: []string{"home", "parent", "people", "human", "child"},
		Name:     "family_man_boy",
		Category: "people",
	},
	"👨‍👧": &emojiMeta{
		Keywords: []string{"home", "parent", "people", "human", "child"},
		Name:     "family_man_girl",
		Category: "people",
	},
	"👨‍👧‍👦": &emojiMeta{
		Keywords: []string{"home", "parent", "people", "human", "children"},
		Name:     "family_man_girl_boy",
		Category: "people",
	},
	"👨‍👦‍👦": &emojiMeta{
		Keywords: []string{"home", "parent", "people", "human", "children"},
		Name:     "family_man_boy_boy",
		Category: "people",
	},
	"👨‍👧‍👧": &emojiMeta{
		Keywords: []string{"home", "parent", "people", "human", "children"},
		Name:     "family_man_girl_girl",
		Category: "people",
	},
	"🧶": &emojiMeta{
		Keywords: []string{"ball", "crochet", "knit"},
		Name:     "yarn",
		Category: "people",
	},
	"🧵": &emojiMeta{
		Keywords: []string{"needle", "sewing", "spool", "string"},
		Name:     "thread",
		Category: "people",
	},
	"🧥": &emojiMeta{
		Keywords: []string{"jacket"},
		Name:     "coat",
		Category: "people",
	},
	"🥼": &emojiMeta{
		Keywords: []string{"doctor", "experiment", "scientist", "chemist"},
		Name:     "labcoat",
		Category: "people",
	},
	"👚": &emojiMeta{
		Keywords: []string{"fashion", "shopping_bags", "female"},
		Name:     "womans_clothes",
		Category: "people",
	},
	"👕": &emojiMeta{
		Keywords: []string{"fashion", "cloth", "casual", "shirt", "tee"},
		Name:     "tshirt",
		Category: "people",
	},
	"👖": &emojiMeta{
		Keywords: []string{"fashion", "shopping"},
		Name:     "jeans",
		Category: "people",
	},
	"👔": &emojiMeta{
		Keywords: []string{"shirt", "suitup", "formal", "fashion", "cloth", "business"},
		Name:     "necktie",
		Category: "people",
	},
	"👗": &emojiMeta{
		Keywords: []string{"clothes", "fashion", "shopping"},
		Name:     "dress",
		Category: "people",
	},
	"👙": &emojiMeta{
		Keywords: []string{"swimming", "female", "woman", "girl", "fashion", "beach", "summer"},
		Name:     "bikini",
		Category: "people",
	},
	"👘": &emojiMeta{
		Keywords: []string{"dress", "fashion", "women", "female", "japanese"},
		Name:     "kimono",
		Category: "people",
	},
	"💄": &emojiMeta{
		Keywords: []string{"female", "girl", "fashion", "woman"},
		Name:     "lipstick",
		Category: "people",
	},
	"💋": &emojiMeta{
		Keywords: []string{"face", "lips", "love", "like", "affection", "valentines"},
		Name:     "kiss",
		Category: "people",
	},
	"👣": &emojiMeta{
		Keywords: []string{"feet", "tracking", "walking", "beach"},
		Name:     "footprints",
		Category: "people",
	},
	"🥿": &emojiMeta{
		Keywords: []string{"ballet", "slip-on", "slipper"},
		Name:     "flat_shoe",
		Category: "people",
	},
	"👠": &emojiMeta{
		Keywords: []string{"fashion", "shoes", "female", "pumps", "stiletto"},
		Name:     "high_heel",
		Category: "people",
	},
	"👡": &emojiMeta{
		Keywords: []string{"shoes", "fashion", "flip flops"},
		Name:     "sandal",
		Category: "people",
	},
	"👢": &emojiMeta{
		Keywords: []string{"shoes", "fashion"},
		Name:     "boot",
		Category: "people",
	},
	"👞": &emojiMeta{
		Keywords: []string{"fashion", "male"},
		Name:     "mans_shoe",
		Category: "people",
	},
	"👟": &emojiMeta{
		Keywords: []string{"shoes", "sports", "sneakers"},
		Name:     "athletic_shoe",
		Category: "people",
	},
	"🥾": &emojiMeta{
		Keywords: []string{"backpacking", "camping", "hiking"},
		Name:     "hiking_boot",
		Category: "people",
	},
	"🧦": &emojiMeta{
		Keywords: []string{"stockings", "clothes"},
		Name:     "socks",
		Category: "people",
	},
	"🧤": &emojiMeta{
		Keywords: []string{"hands", "winter", "clothes"},
		Name:     "gloves",
		Category: "people",
	},
	"🧣": &emojiMeta{
		Keywords: []string{"neck", "winter", "clothes"},
		Name:     "scarf",
		Category: "people",
	},
	"👒": &emojiMeta{
		Keywords: []string{"fashion", "accessories", "female", "lady", "spring"},
		Name:     "womans_hat",
		Category: "people",
	},
	"🎩": &emojiMeta{
		Keywords: []string{"magic", "gentleman", "classy", "circus"},
		Name:     "tophat",
		Category: "people",
	},
	"🧢": &emojiMeta{
		Keywords: []string{"cap", "baseball"},
		Name:     "billed_hat",
		Category: "people",
	},
	"⛑": &emojiMeta{
		Keywords: []string{"construction", "build"},
		Name:     "rescue_worker_helmet",
		Category: "people",
	},
	"🎓": &emojiMeta{
		Keywords: []string{"school", "college", "degree", "university", "graduation", "cap", "hat", "legal", "learn", "education"},
		Name:     "mortar_board",
		Category: "people",
	},
	"👑": &emojiMeta{
		Keywords: []string{"king", "kod", "leader", "royalty", "lord"},
		Name:     "crown",
		Category: "people",
	},
	"🎒": &emojiMeta{
		Keywords: []string{"student", "education", "bag", "backpack"},
		Name:     "school_satchel",
		Category: "people",
	},
	"🧳": &emojiMeta{
		Keywords: []string{"packing", "travel"},
		Name:     "luggage",
		Category: "people",
	},
	"👝": &emojiMeta{
		Keywords: []string{"bag", "accessories", "shopping"},
		Name:     "pouch",
		Category: "people",
	},
	"👛": &emojiMeta{
		Keywords: []string{"fashion", "accessories", "money", "sales", "shopping"},
		Name:     "purse",
		Category: "people",
	},
	"👜": &emojiMeta{
		Keywords: []string{"fashion", "accessory", "accessories", "shopping"},
		Name:     "handbag",
		Category: "people",
	},
	"💼": &emojiMeta{
		Keywords: []string{"business", "documents", "work", "law", "legal", "job", "career"},
		Name:     "briefcase",
		Category: "people",
	},
	"👓": &emojiMeta{
		Keywords: []string{"fashion", "accessories", "eyesight", "nerdy", "dork", "geek"},
		Name:     "eyeglasses",
		Category: "people",
	},
	"🕶": &emojiMeta{
		Keywords: []string{"face", "cool", "accessories"},
		Name:     "dark_sunglasses",
		Category: "people",
	},
	"🥽": &emojiMeta{
		Keywords: []string{"eyes", "protection", "safety"},
		Name:     "goggles",
		Category: "people",
	},
	"💍": &emojiMeta{
		Keywords: []string{"wedding", "propose", "marriage", "valentines", "diamond", "fashion", "jewelry", "gem", "engagement"},
		Name:     "ring",
		Category: "people",
	},
	"🌂": &emojiMeta{
		Keywords: []string{"weather", "rain", "drizzle"},
		Name:     "closed_umbrella",
		Category: "people",
	},
	"🐶": &emojiMeta{
		Keywords: []string{"animal", "friend", "nature", "woof", "puppy", "pet", "faithful"},
		Name:     "dog",
		Category: "animals_and_nature",
	},
	"🐱": &emojiMeta{
		Keywords: []string{"animal", "meow", "nature", "pet", "kitten"},
		Name:     "cat",
		Category: "animals_and_nature",
	},
	"🐭": &emojiMeta{
		Keywords: []string{"animal", "nature", "cheese_wedge", "rodent"},
		Name:     "mouse",
		Category: "animals_and_nature",
	},
	"🐹": &emojiMeta{
		Keywords: []string{"animal", "nature"},
		Name:     "hamster",
		Category: "animals_and_nature",
	},
	"🐰": &emojiMeta{
		Keywords: []string{"animal", "nature", "pet", "spring", "magic", "bunny"},
		Name:     "rabbit",
		Category: "animals_and_nature",
	},
	"🦊": &emojiMeta{
		Keywords: []string{"animal", "nature", "face"},
		Name:     "fox_face",
		Category: "animals_and_nature",
	},
	"🐻": &emojiMeta{
		Keywords: []string{"animal", "nature", "wild"},
		Name:     "bear",
		Category: "animals_and_nature",
	},
	"🐼": &emojiMeta{
		Keywords: []string{"animal", "nature", "panda"},
		Name:     "panda_face",
		Category: "animals_and_nature",
	},
	"🐨": &emojiMeta{
		Keywords: []string{"animal", "nature"},
		Name:     "koala",
		Category: "animals_and_nature",
	},
	"🐯": &emojiMeta{
		Keywords: []string{"animal", "cat", "danger", "wild", "nature", "roar"},
		Name:     "tiger",
		Category: "animals_and_nature",
	},
	"🦁": &emojiMeta{
		Keywords: []string{"animal", "nature"},
		Name:     "lion",
		Category: "animals_and_nature",
	},
	"🐮": &emojiMeta{
		Keywords: []string{"beef", "ox", "animal", "nature", "moo", "milk"},
		Name:     "cow",
		Category: "animals_and_nature",
	},
	"🐷": &emojiMeta{
		Keywords: []string{"animal", "oink", "nature"},
		Name:     "pig",
		Category: "animals_and_nature",
	},
	"🐽": &emojiMeta{
		Keywords: []string{"animal", "oink"},
		Name:     "pig_nose",
		Category: "animals_and_nature",
	},
	"🐸": &emojiMeta{
		Keywords: []string{"animal", "nature", "croak", "toad"},
		Name:     "frog",
		Category: "animals_and_nature",
	},
	"🦑": &emojiMeta{
		Keywords: []string{"animal", "nature", "ocean", "sea"},
		Name:     "squid",
		Category: "animals_and_nature",
	},
	"🐙": &emojiMeta{
		Keywords: []string{"animal", "creature", "ocean", "sea", "nature", "beach"},
		Name:     "octopus",
		Category: "animals_and_nature",
	},
	"🦐": &emojiMeta{
		Keywords: []string{"animal", "ocean", "nature", "seafood"},
		Name:     "shrimp",
		Category: "animals_and_nature",
	},
	"🐵": &emojiMeta{
		Keywords: []string{"animal", "nature", "circus"},
		Name:     "monkey_face",
		Category: "animals_and_nature",
	},
	"🦍": &emojiMeta{
		Keywords: []string{"animal", "nature", "circus"},
		Name:     "gorilla",
		Category: "animals_and_nature",
	},
	"🙈": &emojiMeta{
		Keywords: []string{"monkey", "animal", "nature", "haha"},
		Name:     "see_no_evil",
		Category: "animals_and_nature",
	},
	"🙉": &emojiMeta{
		Keywords: []string{"animal", "monkey", "nature"},
		Name:     "hear_no_evil",
		Category: "animals_and_nature",
	},
	"🙊": &emojiMeta{
		Keywords: []string{"monkey", "animal", "nature", "omg"},
		Name:     "speak_no_evil",
		Category: "animals_and_nature",
	},
	"🐒": &emojiMeta{
		Keywords: []string{"animal", "nature", "banana", "circus"},
		Name:     "monkey",
		Category: "animals_and_nature",
	},
	"🐔": &emojiMeta{
		Keywords: []string{"animal", "cluck", "nature", "bird"},
		Name:     "chicken",
		Category: "animals_and_nature",
	},
	"🐧": &emojiMeta{
		Keywords: []string{"animal", "nature"},
		Name:     "penguin",
		Category: "animals_and_nature",
	},
	"🐦": &emojiMeta{
		Keywords: []string{"animal", "nature", "fly", "tweet", "spring"},
		Name:     "bird",
		Category: "animals_and_nature",
	},
	"🐤": &emojiMeta{
		Keywords: []string{"animal", "chicken", "bird"},
		Name:     "baby_chick",
		Category: "animals_and_nature",
	},
	"🐣": &emojiMeta{
		Keywords: []string{"animal", "chicken", "egg", "born", "baby", "bird"},
		Name:     "hatching_chick",
		Category: "animals_and_nature",
	},
	"🐥": &emojiMeta{
		Keywords: []string{"animal", "chicken", "baby", "bird"},
		Name:     "hatched_chick",
		Category: "animals_and_nature",
	},
	"🦆": &emojiMeta{
		Keywords: []string{"animal", "nature", "bird", "mallard"},
		Name:     "duck",
		Category: "animals_and_nature",
	},
	"🦅": &emojiMeta{
		Keywords: []string{"animal", "nature", "bird"},
		Name:     "eagle",
		Category: "animals_and_nature",
	},
	"🦉": &emojiMeta{
		Keywords: []string{"animal", "nature", "bird", "hoot"},
		Name:     "owl",
		Category: "animals_and_nature",
	},
	"🦇": &emojiMeta{
		Keywords: []string{"animal", "nature", "blind", "vampire"},
		Name:     "bat",
		Category: "animals_and_nature",
	},
	"🐺": &emojiMeta{
		Keywords: []string{"animal", "nature", "wild"},
		Name:     "wolf",
		Category: "animals_and_nature",
	},
	"🐗": &emojiMeta{
		Keywords: []string{"animal", "nature"},
		Name:     "boar",
		Category: "animals_and_nature",
	},
	"🐴": &emojiMeta{
		Keywords: []string{"animal", "brown", "nature"},
		Name:     "horse",
		Category: "animals_and_nature",
	},
	"🦄": &emojiMeta{
		Keywords: []string{"animal", "nature", "mystical"},
		Name:     "unicorn",
		Category: "animals_and_nature",
	},
	"🐝": &emojiMeta{
		Keywords: []string{"animal", "insect", "nature", "bug", "spring", "honey"},
		Name:     "honeybee",
		Category: "animals_and_nature",
	},
	"🐛": &emojiMeta{
		Keywords: []string{"animal", "insect", "nature", "worm"},
		Name:     "bug",
		Category: "animals_and_nature",
	},
	"🦋": &emojiMeta{
		Keywords: []string{"animal", "insect", "nature", "caterpillar"},
		Name:     "butterfly",
		Category: "animals_and_nature",
	},
	"🐌": &emojiMeta{
		Keywords: []string{"slow", "animal", "shell"},
		Name:     "snail",
		Category: "animals_and_nature",
	},
	"🐞": &emojiMeta{
		Keywords: []string{"animal", "insect", "nature", "ladybug"},
		Name:     "beetle",
		Category: "animals_and_nature",
	},
	"🐜": &emojiMeta{
		Keywords: []string{"animal", "insect", "nature", "bug"},
		Name:     "ant",
		Category: "animals_and_nature",
	},
	"🦗": &emojiMeta{
		Keywords: []string{"animal", "cricket", "chirp"},
		Name:     "grasshopper",
		Category: "animals_and_nature",
	},
	"🕷": &emojiMeta{
		Keywords: []string{"animal", "arachnid"},
		Name:     "spider",
		Category: "animals_and_nature",
	},
	"🦂": &emojiMeta{
		Keywords: []string{"animal", "arachnid"},
		Name:     "scorpion",
		Category: "animals_and_nature",
	},
	"🦀": &emojiMeta{
		Keywords: []string{"animal", "crustacean"},
		Name:     "crab",
		Category: "animals_and_nature",
	},
	"🐍": &emojiMeta{
		Keywords: []string{"animal", "evil", "nature", "hiss", "python"},
		Name:     "snake",
		Category: "animals_and_nature",
	},
	"🦎": &emojiMeta{
		Keywords: []string{"animal", "nature", "reptile"},
		Name:     "lizard",
		Category: "animals_and_nature",
	},
	"🦖": &emojiMeta{
		Keywords: []string{"animal", "nature", "dinosaur", "tyrannosaurus", "extinct"},
		Name:     "t-rex",
		Category: "animals_and_nature",
	},
	"🦕": &emojiMeta{
		Keywords: []string{"animal", "nature", "dinosaur", "brachiosaurus", "brontosaurus", "diplodocus", "extinct"},
		Name:     "sauropod",
		Category: "animals_and_nature",
	},
	"🐢": &emojiMeta{
		Keywords: []string{"animal", "slow", "nature", "tortoise"},
		Name:     "turtle",
		Category: "animals_and_nature",
	},
	"🐠": &emojiMeta{
		Keywords: []string{"animal", "swim", "ocean", "beach", "nemo"},
		Name:     "tropical_fish",
		Category: "animals_and_nature",
	},
	"🐟": &emojiMeta{
		Keywords: []string{"animal", "food", "nature"},
		Name:     "fish",
		Category: "animals_and_nature",
	},
	"🐡": &emojiMeta{
		Keywords: []string{"animal", "nature", "food", "sea", "ocean"},
		Name:     "blowfish",
		Category: "animals_and_nature",
	},
	"🐬": &emojiMeta{
		Keywords: []string{"animal", "nature", "fish", "sea", "ocean", "flipper", "fins", "beach"},
		Name:     "dolphin",
		Category: "animals_and_nature",
	},
	"🦈": &emojiMeta{
		Keywords: []string{"animal", "nature", "fish", "sea", "ocean", "jaws", "fins", "beach"},
		Name:     "shark",
		Category: "animals_and_nature",
	},
	"🐳": &emojiMeta{
		Keywords: []string{"animal", "nature", "sea", "ocean"},
		Name:     "whale",
		Category: "animals_and_nature",
	},
	"🐋": &emojiMeta{
		Keywords: []string{"animal", "nature", "sea", "ocean"},
		Name:     "whale2",
		Category: "animals_and_nature",
	},
	"🐊": &emojiMeta{
		Keywords: []string{"animal", "nature", "reptile", "lizard", "alligator"},
		Name:     "crocodile",
		Category: "animals_and_nature",
	},
	"🐆": &emojiMeta{
		Keywords: []string{"animal", "nature"},
		Name:     "leopard",
		Category: "animals_and_nature",
	},
	"🦓": &emojiMeta{
		Keywords: []string{"animal", "nature", "stripes", "safari"},
		Name:     "zebra",
		Category: "animals_and_nature",
	},
	"🐅": &emojiMeta{
		Keywords: []string{"animal", "nature", "roar"},
		Name:     "tiger2",
		Category: "animals_and_nature",
	},
	"🐃": &emojiMeta{
		Keywords: []string{"animal", "nature", "ox", "cow"},
		Name:     "water_buffalo",
		Category: "animals_and_nature",
	},
	"🐂": &emojiMeta{
		Keywords: []string{"animal", "cow", "beef"},
		Name:     "ox",
		Category: "animals_and_nature",
	},
	"🐄": &emojiMeta{
		Keywords: []string{"beef", "ox", "animal", "nature", "moo", "milk"},
		Name:     "cow2",
		Category: "animals_and_nature",
	},
	"🦌": &emojiMeta{
		Keywords: []string{"animal", "nature", "horns", "venison"},
		Name:     "deer",
		Category: "animals_and_nature",
	},
	"🐪": &emojiMeta{
		Keywords: []string{"animal", "hot", "desert", "hump"},
		Name:     "dromedary_camel",
		Category: "animals_and_nature",
	},
	"🐫": &emojiMeta{
		Keywords: []string{"animal", "nature", "hot", "desert", "hump"},
		Name:     "camel",
		Category: "animals_and_nature",
	},
	"🦒": &emojiMeta{
		Keywords: []string{"animal", "nature", "spots", "safari"},
		Name:     "giraffe",
		Category: "animals_and_nature",
	},
	"🐘": &emojiMeta{
		Keywords: []string{"animal", "nature", "nose", "th", "circus"},
		Name:     "elephant",
		Category: "animals_and_nature",
	},
	"🦏": &emojiMeta{
		Keywords: []string{"animal", "nature", "horn"},
		Name:     "rhinoceros",
		Category: "animals_and_nature",
	},
	"🐐": &emojiMeta{
		Keywords: []string{"animal", "nature"},
		Name:     "goat",
		Category: "animals_and_nature",
	},
	"🐏": &emojiMeta{
		Keywords: []string{"animal", "sheep", "nature"},
		Name:     "ram",
		Category: "animals_and_nature",
	},
	"🐑": &emojiMeta{
		Keywords: []string{"animal", "nature", "wool", "shipit"},
		Name:     "sheep",
		Category: "animals_and_nature",
	},
	"🐎": &emojiMeta{
		Keywords: []string{"animal", "gamble", "luck"},
		Name:     "racehorse",
		Category: "animals_and_nature",
	},
	"🐖": &emojiMeta{
		Keywords: []string{"animal", "nature"},
		Name:     "pig2",
		Category: "animals_and_nature",
	},
	"🐀": &emojiMeta{
		Keywords: []string{"animal", "mouse", "rodent"},
		Name:     "rat",
		Category: "animals_and_nature",
	},
	"🐁": &emojiMeta{
		Keywords: []string{"animal", "nature", "rodent"},
		Name:     "mouse2",
		Category: "animals_and_nature",
	},
	"🐓": &emojiMeta{
		Keywords: []string{"animal", "nature", "chicken"},
		Name:     "rooster",
		Category: "animals_and_nature",
	},
	"🦃": &emojiMeta{
		Keywords: []string{"animal", "bird"},
		Name:     "turkey",
		Category: "animals_and_nature",
	},
	"🕊": &emojiMeta{
		Keywords: []string{"animal", "bird"},
		Name:     "dove",
		Category: "animals_and_nature",
	},
	"🐕": &emojiMeta{
		Keywords: []string{"animal", "nature", "friend", "doge", "pet", "faithful"},
		Name:     "dog2",
		Category: "animals_and_nature",
	},
	"🐩": &emojiMeta{
		Keywords: []string{"dog", "animal", "101", "nature", "pet"},
		Name:     "poodle",
		Category: "animals_and_nature",
	},
	"🐈": &emojiMeta{
		Keywords: []string{"animal", "meow", "pet", "cats"},
		Name:     "cat2",
		Category: "animals_and_nature",
	},
	"🐇": &emojiMeta{
		Keywords: []string{"animal", "nature", "pet", "magic", "spring"},
		Name:     "rabbit2",
		Category: "animals_and_nature",
	},
	"🐿": &emojiMeta{
		Keywords: []string{"animal", "nature", "rodent", "squirrel"},
		Name:     "chipmunk",
		Category: "animals_and_nature",
	},
	"🦔": &emojiMeta{
		Keywords: []string{"animal", "nature", "spiny"},
		Name:     "hedgehog",
		Category: "animals_and_nature",
	},
	"🦝": &emojiMeta{
		Keywords: []string{"animal", "nature"},
		Name:     "raccoon",
		Category: "animals_and_nature",
	},
	"🦙": &emojiMeta{
		Keywords: []string{"animal", "nature", "alpaca"},
		Name:     "llama",
		Category: "animals_and_nature",
	},
	"🦛": &emojiMeta{
		Keywords: []string{"animal", "nature"},
		Name:     "hippopotamus",
		Category: "animals_and_nature",
	},
	"🦘": &emojiMeta{
		Keywords: []string{"animal", "nature", "australia", "joey", "hop", "marsupial"},
		Name:     "kangaroo",
		Category: "animals_and_nature",
	},
	"🦡": &emojiMeta{
		Keywords: []string{"animal", "nature", "honey"},
		Name:     "badger",
		Category: "animals_and_nature",
	},
	"🦢": &emojiMeta{
		Keywords: []string{"animal", "nature", "bird"},
		Name:     "swan",
		Category: "animals_and_nature",
	},
	"🦚": &emojiMeta{
		Keywords: []string{"animal", "nature", "peahen", "bird"},
		Name:     "peacock",
		Category: "animals_and_nature",
	},
	"🦜": &emojiMeta{
		Keywords: []string{"animal", "nature", "bird", "pirate", "talk"},
		Name:     "parrot",
		Category: "animals_and_nature",
	},
	"🦞": &emojiMeta{
		Keywords: []string{"animal", "nature", "bisque", "claws", "seafood"},
		Name:     "lobster",
		Category: "animals_and_nature",
	},
	"🦟": &emojiMeta{
		Keywords: []string{"animal", "nature", "insect", "malaria"},
		Name:     "mosquito",
		Category: "animals_and_nature",
	},
	"🐾": &emojiMeta{
		Keywords: []string{"animal", "tracking", "footprints", "dog", "cat", "pet", "feet"},
		Name:     "paw_prints",
		Category: "animals_and_nature",
	},
	"🐉": &emojiMeta{
		Keywords: []string{"animal", "myth", "nature", "chinese", "green"},
		Name:     "dragon",
		Category: "animals_and_nature",
	},
	"🐲": &emojiMeta{
		Keywords: []string{"animal", "myth", "nature", "chinese", "green"},
		Name:     "dragon_face",
		Category: "animals_and_nature",
	},
	"🌵": &emojiMeta{
		Keywords: []string{"vegetable", "plant", "nature"},
		Name:     "cactus",
		Category: "animals_and_nature",
	},
	"🎄": &emojiMeta{
		Keywords: []string{"festival", "vacation", "december", "xmas", "celebration"},
		Name:     "christmas_tree",
		Category: "animals_and_nature",
	},
	"🌲": &emojiMeta{
		Keywords: []string{"plant", "nature"},
		Name:     "evergreen_tree",
		Category: "animals_and_nature",
	},
	"🌳": &emojiMeta{
		Keywords: []string{"plant", "nature"},
		Name:     "deciduous_tree",
		Category: "animals_and_nature",
	},
	"🌴": &emojiMeta{
		Keywords: []string{"plant", "vegetable", "nature", "summer", "beach", "mojito", "tropical"},
		Name:     "palm_tree",
		Category: "animals_and_nature",
	},
	"🌱": &emojiMeta{
		Keywords: []string{"plant", "nature", "grass", "lawn", "spring"},
		Name:     "seedling",
		Category: "animals_and_nature",
	},
	"🌿": &emojiMeta{
		Keywords: []string{"vegetable", "plant", "medicine", "weed", "grass", "lawn"},
		Name:     "herb",
		Category: "animals_and_nature",
	},
	"☘": &emojiMeta{
		Keywords: []string{"vegetable", "plant", "nature", "irish", "clover"},
		Name:     "shamrock",
		Category: "animals_and_nature",
	},
	"🍀": &emojiMeta{
		Keywords: []string{"vegetable", "plant", "nature", "lucky", "irish"},
		Name:     "four_leaf_clover",
		Category: "animals_and_nature",
	},
	"🎍": &emojiMeta{
		Keywords: []string{"plant", "nature", "vegetable", "panda", "pine_decoration"},
		Name:     "bamboo",
		Category: "animals_and_nature",
	},
	"🎋": &emojiMeta{
		Keywords: []string{"plant", "nature", "branch", "summer"},
		Name:     "tanabata_tree",
		Category: "animals_and_nature",
	},
	"🍃": &emojiMeta{
		Keywords: []string{"nature", "plant", "tree", "vegetable", "grass", "lawn", "spring"},
		Name:     "leaves",
		Category: "animals_and_nature",
	},
	"🍂": &emojiMeta{
		Keywords: []string{"nature", "plant", "vegetable", "leaves"},
		Name:     "fallen_leaf",
		Category: "animals_and_nature",
	},
	"🍁": &emojiMeta{
		Keywords: []string{"nature", "plant", "vegetable", "ca", "fall"},
		Name:     "maple_leaf",
		Category: "animals_and_nature",
	},
	"🌾": &emojiMeta{
		Keywords: []string{"nature", "plant"},
		Name:     "ear_of_rice",
		Category: "animals_and_nature",
	},
	"🌺": &emojiMeta{
		Keywords: []string{"plant", "vegetable", "flowers", "beach"},
		Name:     "hibiscus",
		Category: "animals_and_nature",
	},
	"🌻": &emojiMeta{
		Keywords: []string{"nature", "plant", "fall"},
		Name:     "sunflower",
		Category: "animals_and_nature",
	},
	"🌹": &emojiMeta{
		Keywords: []string{"flowers", "valentines", "love", "spring"},
		Name:     "rose",
		Category: "animals_and_nature",
	},
	"🥀": &emojiMeta{
		Keywords: []string{"plant", "nature", "flower"},
		Name:     "wilted_flower",
		Category: "animals_and_nature",
	},
	"🌷": &emojiMeta{
		Keywords: []string{"flowers", "plant", "nature", "summer", "spring"},
		Name:     "tulip",
		Category: "animals_and_nature",
	},
	"🌼": &emojiMeta{
		Keywords: []string{"nature", "flowers", "yellow"},
		Name:     "blossom",
		Category: "animals_and_nature",
	},
	"🌸": &emojiMeta{
		Keywords: []string{"nature", "plant", "spring", "flower"},
		Name:     "cherry_blossom",
		Category: "animals_and_nature",
	},
	"💐": &emojiMeta{
		Keywords: []string{"flowers", "nature", "spring"},
		Name:     "bouquet",
		Category: "animals_and_nature",
	},
	"🍄": &emojiMeta{
		Keywords: []string{"plant", "vegetable"},
		Name:     "mushroom",
		Category: "animals_and_nature",
	},
	"🌰": &emojiMeta{
		Keywords: []string{"food", "squirrel"},
		Name:     "chestnut",
		Category: "animals_and_nature",
	},
	"🎃": &emojiMeta{
		Keywords: []string{"halloween", "light", "pumpkin", "creepy", "fall"},
		Name:     "jack_o_lantern",
		Category: "animals_and_nature",
	},
	"🐚": &emojiMeta{
		Keywords: []string{"nature", "sea", "beach"},
		Name:     "shell",
		Category: "animals_and_nature",
	},
	"🕸": &emojiMeta{
		Keywords: []string{"animal", "insect", "arachnid", "silk"},
		Name:     "spider_web",
		Category: "animals_and_nature",
	},
	"🌎": &emojiMeta{
		Keywords: []string{"globe", "world", "USA", "international"},
		Name:     "earth_americas",
		Category: "animals_and_nature",
	},
	"🌍": &emojiMeta{
		Keywords: []string{"globe", "world", "international"},
		Name:     "earth_africa",
		Category: "animals_and_nature",
	},
	"🌏": &emojiMeta{
		Keywords: []string{"globe", "world", "east", "international"},
		Name:     "earth_asia",
		Category: "animals_and_nature",
	},
	"🌕": &emojiMeta{
		Keywords: []string{"nature", "yellow", "twilight", "planet", "space", "night", "evening", "sleep"},
		Name:     "full_moon",
		Category: "animals_and_nature",
	},
	"🌖": &emojiMeta{
		Keywords: []string{"nature", "twilight", "planet", "space", "night", "evening", "sleep", "waxing_gibbous_moon"},
		Name:     "waning_gibbous_moon",
		Category: "animals_and_nature",
	},
	"🌗": &emojiMeta{
		Keywords: []string{"nature", "twilight", "planet", "space", "night", "evening", "sleep"},
		Name:     "last_quarter_moon",
		Category: "animals_and_nature",
	},
	"🌘": &emojiMeta{
		Keywords: []string{"nature", "twilight", "planet", "space", "night", "evening", "sleep"},
		Name:     "waning_crescent_moon",
		Category: "animals_and_nature",
	},
	"🌑": &emojiMeta{
		Keywords: []string{"nature", "twilight", "planet", "space", "night", "evening", "sleep"},
		Name:     "new_moon",
		Category: "animals_and_nature",
	},
	"🌒": &emojiMeta{
		Keywords: []string{"nature", "twilight", "planet", "space", "night", "evening", "sleep"},
		Name:     "waxing_crescent_moon",
		Category: "animals_and_nature",
	},
	"🌓": &emojiMeta{
		Keywords: []string{"nature", "twilight", "planet", "space", "night", "evening", "sleep"},
		Name:     "first_quarter_moon",
		Category: "animals_and_nature",
	},
	"🌔": &emojiMeta{
		Keywords: []string{"nature", "night", "sky", "gray", "twilight", "planet", "space", "evening", "sleep"},
		Name:     "waxing_gibbous_moon",
		Category: "animals_and_nature",
	},
	"🌚": &emojiMeta{
		Keywords: []string{"nature", "twilight", "planet", "space", "night", "evening", "sleep"},
		Name:     "new_moon_with_face",
		Category: "animals_and_nature",
	},
	"🌝": &emojiMeta{
		Keywords: []string{"nature", "twilight", "planet", "space", "night", "evening", "sleep"},
		Name:     "full_moon_with_face",
		Category: "animals_and_nature",
	},
	"🌛": &emojiMeta{
		Keywords: []string{"nature", "twilight", "planet", "space", "night", "evening", "sleep"},
		Name:     "first_quarter_moon_with_face",
		Category: "animals_and_nature",
	},
	"🌜": &emojiMeta{
		Keywords: []string{"nature", "twilight", "planet", "space", "night", "evening", "sleep"},
		Name:     "last_quarter_moon_with_face",
		Category: "animals_and_nature",
	},
	"🌞": &emojiMeta{
		Keywords: []string{"nature", "morning", "sky"},
		Name:     "sun_with_face",
		Category: "animals_and_nature",
	},
	"🌙": &emojiMeta{
		Keywords: []string{"night", "sleep", "sky", "evening", "magic"},
		Name:     "crescent_moon",
		Category: "animals_and_nature",
	},
	"⭐": &emojiMeta{
		Keywords: []string{"night", "yellow"},
		Name:     "star",
		Category: "animals_and_nature",
	},
	"🌟": &emojiMeta{
		Keywords: []string{"night", "sparkle", "awesome", "good", "magic"},
		Name:     "star2",
		Category: "animals_and_nature",
	},
	"💫": &emojiMeta{
		Keywords: []string{"star", "sparkle", "shoot", "magic"},
		Name:     "dizzy",
		Category: "animals_and_nature",
	},
	"✨": &emojiMeta{
		Keywords: []string{"stars", "shine", "shiny", "cool", "awesome", "good", "magic"},
		Name:     "sparkles",
		Category: "animals_and_nature",
	},
	"☄": &emojiMeta{
		Keywords: []string{"space"},
		Name:     "comet",
		Category: "animals_and_nature",
	},
	"☀️": &emojiMeta{
		Keywords: []string{"weather", "nature", "brightness", "summer", "beach", "spring"},
		Name:     "sunny",
		Category: "animals_and_nature",
	},
	"🌤": &emojiMeta{
		Keywords: []string{"weather"},
		Name:     "sun_behind_small_cloud",
		Category: "animals_and_nature",
	},
	"⛅": &emojiMeta{
		Keywords: []string{"weather", "nature", "cloudy", "morning", "fall", "spring"},
		Name:     "partly_sunny",
		Category: "animals_and_nature",
	},
	"🌥": &emojiMeta{
		Keywords: []string{"weather"},
		Name:     "sun_behind_large_cloud",
		Category: "animals_and_nature",
	},
	"🌦": &emojiMeta{
		Keywords: []string{"weather"},
		Name:     "sun_behind_rain_cloud",
		Category: "animals_and_nature",
	},
	"☁️": &emojiMeta{
		Keywords: []string{"weather", "sky"},
		Name:     "cloud",
		Category: "animals_and_nature",
	},
	"🌧": &emojiMeta{
		Keywords: []string{"weather"},
		Name:     "cloud_with_rain",
		Category: "animals_and_nature",
	},
	"⛈": &emojiMeta{
		Keywords: []string{"weather", "lightning"},
		Name:     "cloud_with_lightning_and_rain",
		Category: "animals_and_nature",
	},
	"🌩": &emojiMeta{
		Keywords: []string{"weather", "thunder"},
		Name:     "cloud_with_lightning",
		Category: "animals_and_nature",
	},
	"⚡": &emojiMeta{
		Keywords: []string{"thunder", "weather", "lightning bolt", "fast"},
		Name:     "zap",
		Category: "animals_and_nature",
	},
	"🔥": &emojiMeta{
		Keywords: []string{"hot", "cook", "flame"},
		Name:     "fire",
		Category: "animals_and_nature",
	},
	"💥": &emojiMeta{
		Keywords: []string{"bomb", "explode", "explosion", "collision", "blown"},
		Name:     "boom",
		Category: "animals_and_nature",
	},
	"❄️": &emojiMeta{
		Keywords: []string{"winter", "season", "cold", "weather", "christmas", "xmas"},
		Name:     "snowflake",
		Category: "animals_and_nature",
	},
	"🌨": &emojiMeta{
		Keywords: []string{"weather"},
		Name:     "cloud_with_snow",
		Category: "animals_and_nature",
	},
	"⛄": &emojiMeta{
		Keywords: []string{"winter", "season", "cold", "weather", "christmas", "xmas", "frozen", "without_snow"},
		Name:     "snowman",
		Category: "animals_and_nature",
	},
	"☃": &emojiMeta{
		Keywords: []string{"winter", "season", "cold", "weather", "christmas", "xmas", "frozen"},
		Name:     "snowman_with_snow",
		Category: "animals_and_nature",
	},
	"🌬": &emojiMeta{
		Keywords: []string{"gust", "air"},
		Name:     "wind_face",
		Category: "animals_and_nature",
	},
	"💨": &emojiMeta{
		Keywords: []string{"wind", "air", "fast", "shoo", "fart", "smoke", "puff"},
		Name:     "dash",
		Category: "animals_and_nature",
	},
	"🌪": &emojiMeta{
		Keywords: []string{"weather", "cyclone", "twister"},
		Name:     "tornado",
		Category: "animals_and_nature",
	},
	"🌫": &emojiMeta{
		Keywords: []string{"weather"},
		Name:     "fog",
		Category: "animals_and_nature",
	},
	"☂": &emojiMeta{
		Keywords: []string{"weather", "spring"},
		Name:     "open_umbrella",
		Category: "animals_and_nature",
	},
	"☔": &emojiMeta{
		Keywords: []string{"rainy", "weather", "spring"},
		Name:     "umbrella",
		Category: "animals_and_nature",
	},
	"💧": &emojiMeta{
		Keywords: []string{"water", "drip", "faucet", "spring"},
		Name:     "droplet",
		Category: "animals_and_nature",
	},
	"💦": &emojiMeta{
		Keywords: []string{"water", "drip", "oops"},
		Name:     "sweat_drops",
		Category: "animals_and_nature",
	},
	"🌊": &emojiMeta{
		Keywords: []string{"sea", "water", "wave", "nature", "tsunami", "disaster"},
		Name:     "ocean",
		Category: "animals_and_nature",
	},
	"🍏": &emojiMeta{
		Keywords: []string{"fruit", "nature"},
		Name:     "green_apple",
		Category: "food_and_drink",
	},
	"🍎": &emojiMeta{
		Keywords: []string{"fruit", "mac", "school"},
		Name:     "apple",
		Category: "food_and_drink",
	},
	"🍐": &emojiMeta{
		Keywords: []string{"fruit", "nature", "food"},
		Name:     "pear",
		Category: "food_and_drink",
	},
	"🍊": &emojiMeta{
		Keywords: []string{"food", "fruit", "nature", "orange"},
		Name:     "tangerine",
		Category: "food_and_drink",
	},
	"🍋": &emojiMeta{
		Keywords: []string{"fruit", "nature"},
		Name:     "lemon",
		Category: "food_and_drink",
	},
	"🍌": &emojiMeta{
		Keywords: []string{"fruit", "food", "monkey"},
		Name:     "banana",
		Category: "food_and_drink",
	},
	"🍉": &emojiMeta{
		Keywords: []string{"fruit", "food", "picnic", "summer"},
		Name:     "watermelon",
		Category: "food_and_drink",
	},
	"🍇": &emojiMeta{
		Keywords: []string{"fruit", "food", "wine"},
		Name:     "grapes",
		Category: "food_and_drink",
	},
	"🍓": &emojiMeta{
		Keywords: []string{"fruit", "food", "nature"},
		Name:     "strawberry",
		Category: "food_and_drink",
	},
	"🍈": &emojiMeta{
		Keywords: []string{"fruit", "nature", "food"},
		Name:     "melon",
		Category: "food_and_drink",
	},
	"🍒": &emojiMeta{
		Keywords: []string{"food", "fruit"},
		Name:     "cherries",
		Category: "food_and_drink",
	},
	"🍑": &emojiMeta{
		Keywords: []string{"fruit", "nature", "food"},
		Name:     "peach",
		Category: "food_and_drink",
	},
	"🍍": &emojiMeta{
		Keywords: []string{"fruit", "nature", "food"},
		Name:     "pineapple",
		Category: "food_and_drink",
	},
	"🥥": &emojiMeta{
		Keywords: []string{"fruit", "nature", "food", "palm"},
		Name:     "coconut",
		Category: "food_and_drink",
	},
	"🥝": &emojiMeta{
		Keywords: []string{"fruit", "food"},
		Name:     "kiwi_fruit",
		Category: "food_and_drink",
	},
	"🥭": &emojiMeta{
		Keywords: []string{"fruit", "food", "tropical"},
		Name:     "mango",
		Category: "food_and_drink",
	},
	"🥑": &emojiMeta{
		Keywords: []string{"fruit", "food"},
		Name:     "avocado",
		Category: "food_and_drink",
	},
	"🥦": &emojiMeta{
		Keywords: []string{"fruit", "food", "vegetable"},
		Name:     "broccoli",
		Category: "food_and_drink",
	},
	"🍅": &emojiMeta{
		Keywords: []string{"fruit", "vegetable", "nature", "food"},
		Name:     "tomato",
		Category: "food_and_drink",
	},
	"🍆": &emojiMeta{
		Keywords: []string{"vegetable", "nature", "food", "aubergine"},
		Name:     "eggplant",
		Category: "food_and_drink",
	},
	"🥒": &emojiMeta{
		Keywords: []string{"fruit", "food", "pickle"},
		Name:     "cucumber",
		Category: "food_and_drink",
	},
	"🥕": &emojiMeta{
		Keywords: []string{"vegetable", "food", "orange"},
		Name:     "carrot",
		Category: "food_and_drink",
	},
	"🌶": &emojiMeta{
		Keywords: []string{"food", "spicy", "chilli", "chili"},
		Name:     "hot_pepper",
		Category: "food_and_drink",
	},
	"🥔": &emojiMeta{
		Keywords: []string{"food", "tuber", "vegatable", "starch"},
		Name:     "potato",
		Category: "food_and_drink",
	},
	"🌽": &emojiMeta{
		Keywords: []string{"food", "vegetable", "plant"},
		Name:     "corn",
		Category: "food_and_drink",
	},
	"🥬": &emojiMeta{
		Keywords: []string{"food", "vegetable", "plant", "bok choy", "cabbage", "kale", "lettuce"},
		Name:     "leafy_greens",
		Category: "food_and_drink",
	},
	"🍠": &emojiMeta{
		Keywords: []string{"food", "nature"},
		Name:     "sweet_potato",
		Category: "food_and_drink",
	},
	"🥜": &emojiMeta{
		Keywords: []string{"food", "nut"},
		Name:     "peanuts",
		Category: "food_and_drink",
	},
	"🍯": &emojiMeta{
		Keywords: []string{"bees", "sweet", "kitchen"},
		Name:     "honey_pot",
		Category: "food_and_drink",
	},
	"🥐": &emojiMeta{
		Keywords: []string{"food", "bread", "french"},
		Name:     "croissant",
		Category: "food_and_drink",
	},
	"🍞": &emojiMeta{
		Keywords: []string{"food", "wheat", "breakfast", "toast"},
		Name:     "bread",
		Category: "food_and_drink",
	},
	"🥖": &emojiMeta{
		Keywords: []string{"food", "bread", "french"},
		Name:     "baguette_bread",
		Category: "food_and_drink",
	},
	"🥯": &emojiMeta{
		Keywords: []string{"food", "bread", "bakery", "schmear"},
		Name:     "bagel",
		Category: "food_and_drink",
	},
	"🥨": &emojiMeta{
		Keywords: []string{"food", "bread", "twisted"},
		Name:     "pretzel",
		Category: "food_and_drink",
	},
	"🧀": &emojiMeta{
		Keywords: []string{"food", "chadder"},
		Name:     "cheese",
		Category: "food_and_drink",
	},
	"🥚": &emojiMeta{
		Keywords: []string{"food", "chicken", "breakfast"},
		Name:     "egg",
		Category: "food_and_drink",
	},
	"🥓": &emojiMeta{
		Keywords: []string{"food", "breakfast", "pork", "pig", "meat"},
		Name:     "bacon",
		Category: "food_and_drink",
	},
	"🥩": &emojiMeta{
		Keywords: []string{"food", "cow", "meat", "cut", "chop", "lambchop", "porkchop"},
		Name:     "steak",
		Category: "food_and_drink",
	},
	"🥞": &emojiMeta{
		Keywords: []string{"food", "breakfast", "flapjacks", "hotcakes"},
		Name:     "pancakes",
		Category: "food_and_drink",
	},
	"🍗": &emojiMeta{
		Keywords: []string{"food", "meat", "drumstick", "bird", "chicken", "turkey"},
		Name:     "poultry_leg",
		Category: "food_and_drink",
	},
	"🍖": &emojiMeta{
		Keywords: []string{"good", "food", "drumstick"},
		Name:     "meat_on_bone",
		Category: "food_and_drink",
	},
	"🦴": &emojiMeta{
		Keywords: []string{"skeleton"},
		Name:     "bone",
		Category: "food_and_drink",
	},
	"🍤": &emojiMeta{
		Keywords: []string{"food", "animal", "appetizer", "summer"},
		Name:     "fried_shrimp",
		Category: "food_and_drink",
	},
	"🍳": &emojiMeta{
		Keywords: []string{"food", "breakfast", "kitchen", "egg"},
		Name:     "fried_egg",
		Category: "food_and_drink",
	},
	"🍔": &emojiMeta{
		Keywords: []string{"meat", "fast food", "beef", "cheeseburger", "mcdonalds", "burger king"},
		Name:     "hamburger",
		Category: "food_and_drink",
	},
	"🍟": &emojiMeta{
		Keywords: []string{"chips", "snack", "fast food"},
		Name:     "fries",
		Category: "food_and_drink",
	},
	"🥙": &emojiMeta{
		Keywords: []string{"food", "flatbread", "stuffed", "gyro"},
		Name:     "stuffed_flatbread",
		Category: "food_and_drink",
	},
	"🌭": &emojiMeta{
		Keywords: []string{"food", "frankfurter"},
		Name:     "hotdog",
		Category: "food_and_drink",
	},
	"🍕": &emojiMeta{
		Keywords: []string{"food", "party"},
		Name:     "pizza",
		Category: "food_and_drink",
	},
	"🥪": &emojiMeta{
		Keywords: []string{"food", "lunch", "bread"},
		Name:     "sandwich",
		Category: "food_and_drink",
	},
	"🥫": &emojiMeta{
		Keywords: []string{"food", "soup"},
		Name:     "canned_food",
		Category: "food_and_drink",
	},
	"🍝": &emojiMeta{
		Keywords: []string{"food", "italian", "noodle"},
		Name:     "spaghetti",
		Category: "food_and_drink",
	},
	"🌮": &emojiMeta{
		Keywords: []string{"food", "mexican"},
		Name:     "taco",
		Category: "food_and_drink",
	},
	"🌯": &emojiMeta{
		Keywords: []string{"food", "mexican"},
		Name:     "burrito",
		Category: "food_and_drink",
	},
	"🥗": &emojiMeta{
		Keywords: []string{"food", "healthy", "lettuce"},
		Name:     "green_salad",
		Category: "food_and_drink",
	},
	"🥘": &emojiMeta{
		Keywords: []string{"food", "cooking", "casserole", "paella"},
		Name:     "shallow_pan_of_food",
		Category: "food_and_drink",
	},
	"🍜": &emojiMeta{
		Keywords: []string{"food", "japanese", "noodle", "chopsticks"},
		Name:     "ramen",
		Category: "food_and_drink",
	},
	"🍲": &emojiMeta{
		Keywords: []string{"food", "meat", "soup"},
		Name:     "stew",
		Category: "food_and_drink",
	},
	"🍥": &emojiMeta{
		Keywords: []string{"food", "japan", "sea", "beach", "narutomaki", "pink", "swirl", "kamaboko", "surimi", "ramen"},
		Name:     "fish_cake",
		Category: "food_and_drink",
	},
	"🥠": &emojiMeta{
		Keywords: []string{"food", "prophecy"},
		Name:     "fortune_cookie",
		Category: "food_and_drink",
	},
	"🍣": &emojiMeta{
		Keywords: []string{"food", "fish", "japanese", "rice"},
		Name:     "sushi",
		Category: "food_and_drink",
	},
	"🍱": &emojiMeta{
		Keywords: []string{"food", "japanese", "box"},
		Name:     "bento",
		Category: "food_and_drink",
	},
	"🍛": &emojiMeta{
		Keywords: []string{"food", "spicy", "hot", "indian"},
		Name:     "curry",
		Category: "food_and_drink",
	},
	"🍙": &emojiMeta{
		Keywords: []string{"food", "japanese"},
		Name:     "rice_ball",
		Category: "food_and_drink",
	},
	"🍚": &emojiMeta{
		Keywords: []string{"food", "china", "asian"},
		Name:     "rice",
		Category: "food_and_drink",
	},
	"🍘": &emojiMeta{
		Keywords: []string{"food", "japanese"},
		Name:     "rice_cracker",
		Category: "food_and_drink",
	},
	"🍢": &emojiMeta{
		Keywords: []string{"food", "japanese"},
		Name:     "oden",
		Category: "food_and_drink",
	},
	"🍡": &emojiMeta{
		Keywords: []string{"food", "dessert", "sweet", "japanese", "barbecue", "meat"},
		Name:     "dango",
		Category: "food_and_drink",
	},
	"🍧": &emojiMeta{
		Keywords: []string{"hot", "dessert", "summer"},
		Name:     "shaved_ice",
		Category: "food_and_drink",
	},
	"🍨": &emojiMeta{
		Keywords: []string{"food", "hot", "dessert"},
		Name:     "ice_cream",
		Category: "food_and_drink",
	},
	"🍦": &emojiMeta{
		Keywords: []string{"food", "hot", "dessert", "summer"},
		Name:     "icecream",
		Category: "food_and_drink",
	},
	"🥧": &emojiMeta{
		Keywords: []string{"food", "dessert", "pastry"},
		Name:     "pie",
		Category: "food_and_drink",
	},
	"🍰": &emojiMeta{
		Keywords: []string{"food", "dessert"},
		Name:     "cake",
		Category: "food_and_drink",
	},
	"🧁": &emojiMeta{
		Keywords: []string{"food", "dessert", "bakery", "sweet"},
		Name:     "cupcake",
		Category: "food_and_drink",
	},
	"🥮": &emojiMeta{
		Keywords: []string{"food", "autumn"},
		Name:     "moon_cake",
		Category: "food_and_drink",
	},
	"🎂": &emojiMeta{
		Keywords: []string{"food", "dessert", "cake"},
		Name:     "birthday",
		Category: "food_and_drink",
	},
	"🍮": &emojiMeta{
		Keywords: []string{"dessert", "food"},
		Name:     "custard",
		Category: "food_and_drink",
	},
	"🍬": &emojiMeta{
		Keywords: []string{"snack", "dessert", "sweet", "lolly"},
		Name:     "candy",
		Category: "food_and_drink",
	},
	"🍭": &emojiMeta{
		Keywords: []string{"food", "snack", "candy", "sweet"},
		Name:     "lollipop",
		Category: "food_and_drink",
	},
	"🍫": &emojiMeta{
		Keywords: []string{"food", "snack", "dessert", "sweet"},
		Name:     "chocolate_bar",
		Category: "food_and_drink",
	},
	"🍿": &emojiMeta{
		Keywords: []string{"food", "movie theater", "films", "snack"},
		Name:     "popcorn",
		Category: "food_and_drink",
	},
	"🥟": &emojiMeta{
		Keywords: []string{"food", "empanada", "pierogi", "potsticker"},
		Name:     "dumpling",
		Category: "food_and_drink",
	},
	"🍩": &emojiMeta{
		Keywords: []string{"food", "dessert", "snack", "sweet", "donut"},
		Name:     "doughnut",
		Category: "food_and_drink",
	},
	"🍪": &emojiMeta{
		Keywords: []string{"food", "snack", "oreo", "chocolate", "sweet", "dessert"},
		Name:     "cookie",
		Category: "food_and_drink",
	},
	"🥛": &emojiMeta{
		Keywords: []string{"beverage", "drink", "cow"},
		Name:     "milk_glass",
		Category: "food_and_drink",
	},
	"🍺": &emojiMeta{
		Keywords: []string{"relax", "beverage", "drink", "drunk", "party", "pub", "summer", "alcohol", "booze"},
		Name:     "beer",
		Category: "food_and_drink",
	},
	"🍻": &emojiMeta{
		Keywords: []string{"relax", "beverage", "drink", "drunk", "party", "pub", "summer", "alcohol", "booze"},
		Name:     "beers",
		Category: "food_and_drink",
	},
	"🥂": &emojiMeta{
		Keywords: []string{"beverage", "drink", "party", "alcohol", "celebrate", "cheers", "wine", "champagne", "toast"},
		Name:     "clinking_glasses",
		Category: "food_and_drink",
	},
	"🍷": &emojiMeta{
		Keywords: []string{"drink", "beverage", "drunk", "alcohol", "booze"},
		Name:     "wine_glass",
		Category: "food_and_drink",
	},
	"🥃": &emojiMeta{
		Keywords: []string{"drink", "beverage", "drunk", "alcohol", "liquor", "booze", "bourbon", "scotch", "whisky", "glass", "shot"},
		Name:     "tumbler_glass",
		Category: "food_and_drink",
	},
	"🍸": &emojiMeta{
		Keywords: []string{"drink", "drunk", "alcohol", "beverage", "booze", "mojito"},
		Name:     "cocktail",
		Category: "food_and_drink",
	},
	"🍹": &emojiMeta{
		Keywords: []string{"beverage", "cocktail", "summer", "beach", "alcohol", "booze", "mojito"},
		Name:     "tropical_drink",
		Category: "food_and_drink",
	},
	"🍾": &emojiMeta{
		Keywords: []string{"drink", "wine", "bottle", "celebration"},
		Name:     "champagne",
		Category: "food_and_drink",
	},
	"🍶": &emojiMeta{
		Keywords: []string{"wine", "drink", "drunk", "beverage", "japanese", "alcohol", "booze"},
		Name:     "sake",
		Category: "food_and_drink",
	},
	"🍵": &emojiMeta{
		Keywords: []string{"drink", "bowl", "breakfast", "green", "british"},
		Name:     "tea",
		Category: "food_and_drink",
	},
	"🥤": &emojiMeta{
		Keywords: []string{"drink", "soda"},
		Name:     "cup_with_straw",
		Category: "food_and_drink",
	},
	"☕": &emojiMeta{
		Keywords: []string{"beverage", "caffeine", "latte", "espresso"},
		Name:     "coffee",
		Category: "food_and_drink",
	},
	"🍼": &emojiMeta{
		Keywords: []string{"food", "container", "milk"},
		Name:     "baby_bottle",
		Category: "food_and_drink",
	},
	"🧂": &emojiMeta{
		Keywords: []string{"condiment", "shaker"},
		Name:     "salt",
		Category: "food_and_drink",
	},
	"🥄": &emojiMeta{
		Keywords: []string{"cutlery", "kitchen", "tableware"},
		Name:     "spoon",
		Category: "food_and_drink",
	},
	"🍴": &emojiMeta{
		Keywords: []string{"cutlery", "kitchen"},
		Name:     "fork_and_knife",
		Category: "food_and_drink",
	},
	"🍽": &emojiMeta{
		Keywords: []string{"food", "eat", "meal", "lunch", "dinner", "restaurant"},
		Name:     "plate_with_cutlery",
		Category: "food_and_drink",
	},
	"🥣": &emojiMeta{
		Keywords: []string{"food", "breakfast", "cereal", "oatmeal", "porridge"},
		Name:     "bowl_with_spoon",
		Category: "food_and_drink",
	},
	"🥡": &emojiMeta{
		Keywords: []string{"food", "leftovers"},
		Name:     "takeout_box",
		Category: "food_and_drink",
	},
	"🥢": &emojiMeta{
		Keywords: []string{"food"},
		Name:     "chopsticks",
		Category: "food_and_drink",
	},
	"⚽": &emojiMeta{
		Keywords: []string{"sports", "football"},
		Name:     "soccer",
		Category: "activity",
	},
	"🏀": &emojiMeta{
		Keywords: []string{"sports", "balls", "NBA"},
		Name:     "basketball",
		Category: "activity",
	},
	"🏈": &emojiMeta{
		Keywords: []string{"sports", "balls", "NFL"},
		Name:     "football",
		Category: "activity",
	},
	"⚾": &emojiMeta{
		Keywords: []string{"sports", "balls"},
		Name:     "baseball",
		Category: "activity",
	},
	"🥎": &emojiMeta{
		Keywords: []string{"sports", "balls"},
		Name:     "softball",
		Category: "activity",
	},
	"🎾": &emojiMeta{
		Keywords: []string{"sports", "balls", "green"},
		Name:     "tennis",
		Category: "activity",
	},
	"🏐": &emojiMeta{
		Keywords: []string{"sports", "balls"},
		Name:     "volleyball",
		Category: "activity",
	},
	"🏉": &emojiMeta{
		Keywords: []string{"sports", "team"},
		Name:     "rugby_football",
		Category: "activity",
	},
	"🥏": &emojiMeta{
		Keywords: []string{"sports", "frisbee", "ultimate"},
		Name:     "flying_disc",
		Category: "activity",
	},
	"🎱": &emojiMeta{
		Keywords: []string{"pool", "hobby", "game", "luck", "magic"},
		Name:     "8ball",
		Category: "activity",
	},
	"⛳": &emojiMeta{
		Keywords: []string{"sports", "business", "flag", "hole", "summer"},
		Name:     "golf",
		Category: "activity",
	},
	"🏌️‍♀️": &emojiMeta{
		Keywords: []string{"sports", "business", "woman", "female"},
		Name:     "golfing_woman",
		Category: "activity",
	},
	"🏌": &emojiMeta{
		Keywords: []string{"sports", "business"},
		Name:     "golfing_man",
		Category: "activity",
	},
	"🏓": &emojiMeta{
		Keywords: []string{"sports", "pingpong"},
		Name:     "ping_pong",
		Category: "activity",
	},
	"🏸": &emojiMeta{
		Keywords: []string{"sports"},
		Name:     "badminton",
		Category: "activity",
	},
	"🥅": &emojiMeta{
		Keywords: []string{"sports"},
		Name:     "goal_net",
		Category: "activity",
	},
	"🏒": &emojiMeta{
		Keywords: []string{"sports"},
		Name:     "ice_hockey",
		Category: "activity",
	},
	"🏑": &emojiMeta{
		Keywords: []string{"sports"},
		Name:     "field_hockey",
		Category: "activity",
	},
	"🥍": &emojiMeta{
		Keywords: []string{"sports", "ball", "stick"},
		Name:     "lacrosse",
		Category: "activity",
	},
	"🏏": &emojiMeta{
		Keywords: []string{"sports"},
		Name:     "cricket",
		Category: "activity",
	},
	"🎿": &emojiMeta{
		Keywords: []string{"sports", "winter", "cold", "snow"},
		Name:     "ski",
		Category: "activity",
	},
	"⛷": &emojiMeta{
		Keywords: []string{"sports", "winter", "snow"},
		Name:     "skier",
		Category: "activity",
	},
	"🏂": &emojiMeta{
		Keywords: []string{"sports", "winter"},
		Name:     "snowboarder",
		Category: "activity",
	},
	"🤺": &emojiMeta{
		Keywords: []string{"sports", "fencing", "sword"},
		Name:     "person_fencing",
		Category: "activity",
	},
	"🤼‍♀️": &emojiMeta{
		Keywords: []string{"sports", "wrestlers"},
		Name:     "women_wrestling",
		Category: "activity",
	},
	"🤼‍♂️": &emojiMeta{
		Keywords: []string{"sports", "wrestlers"},
		Name:     "men_wrestling",
		Category: "activity",
	},
	"🤸‍♀️": &emojiMeta{
		Keywords: []string{"gymnastics"},
		Name:     "woman_cartwheeling",
		Category: "activity",
	},
	"🤸‍♂️": &emojiMeta{
		Keywords: []string{"gymnastics"},
		Name:     "man_cartwheeling",
		Category: "activity",
	},
	"🤾‍♀️": &emojiMeta{
		Keywords: []string{"sports"},
		Name:     "woman_playing_handball",
		Category: "activity",
	},
	"🤾‍♂️": &emojiMeta{
		Keywords: []string{"sports"},
		Name:     "man_playing_handball",
		Category: "activity",
	},
	"⛸": &emojiMeta{
		Keywords: []string{"sports"},
		Name:     "ice_skate",
		Category: "activity",
	},
	"🥌": &emojiMeta{
		Keywords: []string{"sports"},
		Name:     "curling_stone",
		Category: "activity",
	},
	"🛹": &emojiMeta{
		Keywords: []string{"board"},
		Name:     "skateboard",
		Category: "activity",
	},
	"🛷": &emojiMeta{
		Keywords: []string{"sleigh", "luge", "toboggan"},
		Name:     "sled",
		Category: "activity",
	},
	"🏹": &emojiMeta{
		Keywords: []string{"sports"},
		Name:     "bow_and_arrow",
		Category: "activity",
	},
	"🎣": &emojiMeta{
		Keywords: []string{"food", "hobby", "summer"},
		Name:     "fishing_pole_and_fish",
		Category: "activity",
	},
	"🥊": &emojiMeta{
		Keywords: []string{"sports", "fighting"},
		Name:     "boxing_glove",
		Category: "activity",
	},
	"🥋": &emojiMeta{
		Keywords: []string{"judo", "karate", "taekwondo"},
		Name:     "martial_arts_uniform",
		Category: "activity",
	},
	"🚣‍♀️": &emojiMeta{
		Keywords: []string{"sports", "hobby", "water", "ship", "woman", "female"},
		Name:     "rowing_woman",
		Category: "activity",
	},
	"🚣": &emojiMeta{
		Keywords: []string{"sports", "hobby", "water", "ship"},
		Name:     "rowing_man",
		Category: "activity",
	},
	"🧗‍♀️": &emojiMeta{
		Keywords: []string{"sports", "hobby", "woman", "female", "rock"},
		Name:     "climbing_woman",
		Category: "activity",
	},
	"🧗‍♂️": &emojiMeta{
		Keywords: []string{"sports", "hobby", "man", "male", "rock"},
		Name:     "climbing_man",
		Category: "activity",
	},
	"🏊‍♀️": &emojiMeta{
		Keywords: []string{"sports", "exercise", "human", "athlete", "water", "summer", "woman", "female"},
		Name:     "swimming_woman",
		Category: "activity",
	},
	"🏊": &emojiMeta{
		Keywords: []string{"sports", "exercise", "human", "athlete", "water", "summer"},
		Name:     "swimming_man",
		Category: "activity",
	},
	"🤽‍♀️": &emojiMeta{
		Keywords: []string{"sports", "pool"},
		Name:     "woman_playing_water_polo",
		Category: "activity",
	},
	"🤽‍♂️": &emojiMeta{
		Keywords: []string{"sports", "pool"},
		Name:     "man_playing_water_polo",
		Category: "activity",
	},
	"🧘‍♀️": &emojiMeta{
		Keywords: []string{"woman", "female", "meditation", "yoga", "serenity", "zen", "mindfulness"},
		Name:     "woman_in_lotus_position",
		Category: "activity",
	},
	"🧘‍♂️": &emojiMeta{
		Keywords: []string{"man", "male", "meditation", "yoga", "serenity", "zen", "mindfulness"},
		Name:     "man_in_lotus_position",
		Category: "activity",
	},
	"🏄‍♀️": &emojiMeta{
		Keywords: []string{"sports", "ocean", "sea", "summer", "beach", "woman", "female"},
		Name:     "surfing_woman",
		Category: "activity",
	},
	"🏄": &emojiMeta{
		Keywords: []string{"sports", "ocean", "sea", "summer", "beach"},
		Name:     "surfing_man",
		Category: "activity",
	},
	"🛀": &emojiMeta{
		Keywords: []string{"clean", "shower", "bathroom"},
		Name:     "bath",
		Category: "activity",
	},
	"⛹️‍♀️": &emojiMeta{
		Keywords: []string{"sports", "human", "woman", "female"},
		Name:     "basketball_woman",
		Category: "activity",
	},
	"⛹": &emojiMeta{
		Keywords: []string{"sports", "human"},
		Name:     "basketball_man",
		Category: "activity",
	},
	"🏋️‍♀️": &emojiMeta{
		Keywords: []string{"sports", "training", "exercise", "woman", "female"},
		Name:     "weight_lifting_woman",
		Category: "activity",
	},
	"🏋": &emojiMeta{
		Keywords: []string{"sports", "training", "exercise"},
		Name:     "weight_lifting_man",
		Category: "activity",
	},
	"🚴‍♀️": &emojiMeta{
		Keywords: []string{"sports", "bike", "exercise", "hipster", "woman", "female"},
		Name:     "biking_woman",
		Category: "activity",
	},
	"🚴": &emojiMeta{
		Keywords: []string{"sports", "bike", "exercise", "hipster"},
		Name:     "biking_man",
		Category: "activity",
	},
	"🚵‍♀️": &emojiMeta{
		Keywords: []string{"transportation", "sports", "human", "race", "bike", "woman", "female"},
		Name:     "mountain_biking_woman",
		Category: "activity",
	},
	"🚵": &emojiMeta{
		Keywords: []string{"transportation", "sports", "human", "race", "bike"},
		Name:     "mountain_biking_man",
		Category: "activity",
	},
	"🏇": &emojiMeta{
		Keywords: []string{"animal", "betting", "competition", "gambling", "luck"},
		Name:     "horse_racing",
		Category: "activity",
	},
	"🕴": &emojiMeta{
		Keywords: []string{"suit", "business", "levitate", "hover", "jump"},
		Name:     "business_suit_levitating",
		Category: "activity",
	},
	"🏆": &emojiMeta{
		Keywords: []string{"win", "award", "contest", "place", "ftw", "ceremony"},
		Name:     "trophy",
		Category: "activity",
	},
	"🎽": &emojiMeta{
		Keywords: []string{"play", "pageant"},
		Name:     "running_shirt_with_sash",
		Category: "activity",
	},
	"🏅": &emojiMeta{
		Keywords: []string{"award", "winning"},
		Name:     "medal_sports",
		Category: "activity",
	},
	"🎖": &emojiMeta{
		Keywords: []string{"award", "winning", "army"},
		Name:     "medal_military",
		Category: "activity",
	},
	"🥇": &emojiMeta{
		Keywords: []string{"award", "winning", "first"},
		Name:     "1st_place_medal",
		Category: "activity",
	},
	"🥈": &emojiMeta{
		Keywords: []string{"award", "second"},
		Name:     "2nd_place_medal",
		Category: "activity",
	},
	"🥉": &emojiMeta{
		Keywords: []string{"award", "third"},
		Name:     "3rd_place_medal",
		Category: "activity",
	},
	"🎗": &emojiMeta{
		Keywords: []string{"sports", "cause", "support", "awareness"},
		Name:     "reminder_ribbon",
		Category: "activity",
	},
	"🏵": &emojiMeta{
		Keywords: []string{"flower", "decoration", "military"},
		Name:     "rosette",
		Category: "activity",
	},
	"🎫": &emojiMeta{
		Keywords: []string{"event", "concert", "pass"},
		Name:     "ticket",
		Category: "activity",
	},
	"🎟": &emojiMeta{
		Keywords: []string{"sports", "concert", "entrance"},
		Name:     "tickets",
		Category: "activity",
	},
	"🎭": &emojiMeta{
		Keywords: []string{"acting", "theater", "drama"},
		Name:     "performing_arts",
		Category: "activity",
	},
	"🎨": &emojiMeta{
		Keywords: []string{"design", "paint", "draw", "colors"},
		Name:     "art",
		Category: "activity",
	},
	"🎪": &emojiMeta{
		Keywords: []string{"festival", "carnival", "party"},
		Name:     "circus_tent",
		Category: "activity",
	},
	"🤹‍♀️": &emojiMeta{
		Keywords: []string{"juggle", "balance", "skill", "multitask"},
		Name:     "woman_juggling",
		Category: "activity",
	},
	"🤹‍♂️": &emojiMeta{
		Keywords: []string{"juggle", "balance", "skill", "multitask"},
		Name:     "man_juggling",
		Category: "activity",
	},
	"🎤": &emojiMeta{
		Keywords: []string{"sound", "music", "PA", "sing", "talkshow"},
		Name:     "microphone",
		Category: "activity",
	},
	"🎧": &emojiMeta{
		Keywords: []string{"music", "score", "gadgets"},
		Name:     "headphones",
		Category: "activity",
	},
	"🎼": &emojiMeta{
		Keywords: []string{"treble", "clef", "compose"},
		Name:     "musical_score",
		Category: "activity",
	},
	"🎹": &emojiMeta{
		Keywords: []string{"piano", "instrument", "compose"},
		Name:     "musical_keyboard",
		Category: "activity",
	},
	"🥁": &emojiMeta{
		Keywords: []string{"music", "instrument", "drumsticks", "snare"},
		Name:     "drum",
		Category: "activity",
	},
	"🎷": &emojiMeta{
		Keywords: []string{"music", "instrument", "jazz", "blues"},
		Name:     "saxophone",
		Category: "activity",
	},
	"🎺": &emojiMeta{
		Keywords: []string{"music", "brass"},
		Name:     "trumpet",
		Category: "activity",
	},
	"🎸": &emojiMeta{
		Keywords: []string{"music", "instrument"},
		Name:     "guitar",
		Category: "activity",
	},
	"🎻": &emojiMeta{
		Keywords: []string{"music", "instrument", "orchestra", "symphony"},
		Name:     "violin",
		Category: "activity",
	},
	"🎬": &emojiMeta{
		Keywords: []string{"movie", "film", "record"},
		Name:     "clapper",
		Category: "activity",
	},
	"🎮": &emojiMeta{
		Keywords: []string{"play", "console", "PS4", "controller"},
		Name:     "video_game",
		Category: "activity",
	},
	"👾": &emojiMeta{
		Keywords: []string{"game", "arcade", "play"},
		Name:     "space_invader",
		Category: "activity",
	},
	"🎯": &emojiMeta{
		Keywords: []string{"game", "play", "bar", "target", "bullseye"},
		Name:     "dart",
		Category: "activity",
	},
	"🎲": &emojiMeta{
		Keywords: []string{"dice", "random", "tabletop", "play", "luck"},
		Name:     "game_die",
		Category: "activity",
	},
	"♟": &emojiMeta{
		Keywords: []string{"expendable"},
		Name:     "chess_pawn",
		Category: "activity",
	},
	"🎰": &emojiMeta{
		Keywords: []string{"bet", "gamble", "vegas", "fruit machine", "luck", "casino"},
		Name:     "slot_machine",
		Category: "activity",
	},
	"🧩": &emojiMeta{
		Keywords: []string{"interlocking", "puzzle", "piece"},
		Name:     "jigsaw",
		Category: "activity",
	},
	"🎳": &emojiMeta{
		Keywords: []string{"sports", "fun", "play"},
		Name:     "bowling",
		Category: "activity",
	},
	"🚗": &emojiMeta{
		Keywords: []string{"red", "transportation", "vehicle"},
		Name:     "red_car",
		Category: "travel_and_places",
	},
	"🚕": &emojiMeta{
		Keywords: []string{"uber", "vehicle", "cars", "transportation"},
		Name:     "taxi",
		Category: "travel_and_places",
	},
	"🚙": &emojiMeta{
		Keywords: []string{"transportation", "vehicle"},
		Name:     "blue_car",
		Category: "travel_and_places",
	},
	"🚌": &emojiMeta{
		Keywords: []string{"car", "vehicle", "transportation"},
		Name:     "bus",
		Category: "travel_and_places",
	},
	"🚎": &emojiMeta{
		Keywords: []string{"bart", "transportation", "vehicle"},
		Name:     "trolleybus",
		Category: "travel_and_places",
	},
	"🏎": &emojiMeta{
		Keywords: []string{"sports", "race", "fast", "formula", "f1"},
		Name:     "racing_car",
		Category: "travel_and_places",
	},
	"🚓": &emojiMeta{
		Keywords: []string{"vehicle", "cars", "transportation", "law", "legal", "enforcement"},
		Name:     "police_car",
		Category: "travel_and_places",
	},
	"🚑": &emojiMeta{
		Keywords: []string{"health", "911", "hospital"},
		Name:     "ambulance",
		Category: "travel_and_places",
	},
	"🚒": &emojiMeta{
		Keywords: []string{"transportation", "cars", "vehicle"},
		Name:     "fire_engine",
		Category: "travel_and_places",
	},
	"🚐": &emojiMeta{
		Keywords: []string{"vehicle", "car", "transportation"},
		Name:     "minibus",
		Category: "travel_and_places",
	},
	"🚚": &emojiMeta{
		Keywords: []string{"cars", "transportation"},
		Name:     "truck",
		Category: "travel_and_places",
	},
	"🚛": &emojiMeta{
		Keywords: []string{"vehicle", "cars", "transportation", "express"},
		Name:     "articulated_lorry",
		Category: "travel_and_places",
	},
	"🚜": &emojiMeta{
		Keywords: []string{"vehicle", "car", "farming", "agriculture"},
		Name:     "tractor",
		Category: "travel_and_places",
	},
	"🛴": &emojiMeta{
		Keywords: []string{"vehicle", "kick", "razor"},
		Name:     "kick_scooter",
		Category: "travel_and_places",
	},
	"🏍": &emojiMeta{
		Keywords: []string{"race", "sports", "fast"},
		Name:     "motorcycle",
		Category: "travel_and_places",
	},
	"🚲": &emojiMeta{
		Keywords: []string{"sports", "bicycle", "exercise", "hipster"},
		Name:     "bike",
		Category: "travel_and_places",
	},
	"🛵": &emojiMeta{
		Keywords: []string{"vehicle", "vespa", "sasha"},
		Name:     "motor_scooter",
		Category: "travel_and_places",
	},
	"🚨": &emojiMeta{
		Keywords: []string{"police", "ambulance", "911", "emergency", "alert", "error", "pinged", "law", "legal"},
		Name:     "rotating_light",
		Category: "travel_and_places",
	},
	"🚔": &emojiMeta{
		Keywords: []string{"vehicle", "law", "legal", "enforcement", "911"},
		Name:     "oncoming_police_car",
		Category: "travel_and_places",
	},
	"🚍": &emojiMeta{
		Keywords: []string{"vehicle", "transportation"},
		Name:     "oncoming_bus",
		Category: "travel_and_places",
	},
	"🚘": &emojiMeta{
		Keywords: []string{"car", "vehicle", "transportation"},
		Name:     "oncoming_automobile",
		Category: "travel_and_places",
	},
	"🚖": &emojiMeta{
		Keywords: []string{"vehicle", "cars", "uber"},
		Name:     "oncoming_taxi",
		Category: "travel_and_places",
	},
	"🚡": &emojiMeta{
		Keywords: []string{"transportation", "vehicle", "ski"},
		Name:     "aerial_tramway",
		Category: "travel_and_places",
	},
	"🚠": &emojiMeta{
		Keywords: []string{"transportation", "vehicle", "ski"},
		Name:     "mountain_cableway",
		Category: "travel_and_places",
	},
	"🚟": &emojiMeta{
		Keywords: []string{"vehicle", "transportation"},
		Name:     "suspension_railway",
		Category: "travel_and_places",
	},
	"🚃": &emojiMeta{
		Keywords: []string{"transportation", "vehicle"},
		Name:     "railway_car",
		Category: "travel_and_places",
	},
	"🚋": &emojiMeta{
		Keywords: []string{"transportation", "vehicle", "carriage", "public", "travel"},
		Name:     "train",
		Category: "travel_and_places",
	},
	"🚝": &emojiMeta{
		Keywords: []string{"transportation", "vehicle"},
		Name:     "monorail",
		Category: "travel_and_places",
	},
	"🚄": &emojiMeta{
		Keywords: []string{"transportation", "vehicle"},
		Name:     "bullettrain_side",
		Category: "travel_and_places",
	},
	"🚅": &emojiMeta{
		Keywords: []string{"transportation", "vehicle", "speed", "fast", "public", "travel"},
		Name:     "bullettrain_front",
		Category: "travel_and_places",
	},
	"🚈": &emojiMeta{
		Keywords: []string{"transportation", "vehicle"},
		Name:     "light_rail",
		Category: "travel_and_places",
	},
	"🚞": &emojiMeta{
		Keywords: []string{"transportation", "vehicle"},
		Name:     "mountain_railway",
		Category: "travel_and_places",
	},
	"🚂": &emojiMeta{
		Keywords: []string{"transportation", "vehicle", "train"},
		Name:     "steam_locomotive",
		Category: "travel_and_places",
	},
	"🚆": &emojiMeta{
		Keywords: []string{"transportation", "vehicle"},
		Name:     "train2",
		Category: "travel_and_places",
	},
	"🚇": &emojiMeta{
		Keywords: []string{"transportation", "blue-square", "mrt", "underground", "tube"},
		Name:     "metro",
		Category: "travel_and_places",
	},
	"🚊": &emojiMeta{
		Keywords: []string{"transportation", "vehicle"},
		Name:     "tram",
		Category: "travel_and_places",
	},
	"🚉": &emojiMeta{
		Keywords: []string{"transportation", "vehicle", "public"},
		Name:     "station",
		Category: "travel_and_places",
	},
	"🛸": &emojiMeta{
		Keywords: []string{"transportation", "vehicle", "ufo"},
		Name:     "flying_saucer",
		Category: "travel_and_places",
	},
	"🚁": &emojiMeta{
		Keywords: []string{"transportation", "vehicle", "fly"},
		Name:     "helicopter",
		Category: "travel_and_places",
	},
	"🛩": &emojiMeta{
		Keywords: []string{"flight", "transportation", "fly", "vehicle"},
		Name:     "small_airplane",
		Category: "travel_and_places",
	},
	"✈️": &emojiMeta{
		Keywords: []string{"vehicle", "transportation", "flight", "fly"},
		Name:     "airplane",
		Category: "travel_and_places",
	},
	"🛫": &emojiMeta{
		Keywords: []string{"airport", "flight", "landing"},
		Name:     "flight_departure",
		Category: "travel_and_places",
	},
	"🛬": &emojiMeta{
		Keywords: []string{"airport", "flight", "boarding"},
		Name:     "flight_arrival",
		Category: "travel_and_places",
	},
	"⛵": &emojiMeta{
		Keywords: []string{"ship", "summer", "transportation", "water", "sailing"},
		Name:     "sailboat",
		Category: "travel_and_places",
	},
	"🛥": &emojiMeta{
		Keywords: []string{"ship"},
		Name:     "motor_boat",
		Category: "travel_and_places",
	},
	"🚤": &emojiMeta{
		Keywords: []string{"ship", "transportation", "vehicle", "summer"},
		Name:     "speedboat",
		Category: "travel_and_places",
	},
	"⛴": &emojiMeta{
		Keywords: []string{"boat", "ship", "yacht"},
		Name:     "ferry",
		Category: "travel_and_places",
	},
	"🛳": &emojiMeta{
		Keywords: []string{"yacht", "cruise", "ferry"},
		Name:     "passenger_ship",
		Category: "travel_and_places",
	},
	"🚀": &emojiMeta{
		Keywords: []string{"launch", "ship", "staffmode", "NASA", "outer space", "outer_space", "fly"},
		Name:     "rocket",
		Category: "travel_and_places",
	},
	"🛰": &emojiMeta{
		Keywords: []string{"communication", "gps", "orbit", "spaceflight", "NASA", "ISS"},
		Name:     "artificial_satellite",
		Category: "travel_and_places",
	},
	"💺": &emojiMeta{
		Keywords: []string{"sit", "airplane", "transport", "bus", "flight", "fly"},
		Name:     "seat",
		Category: "travel_and_places",
	},
	"🛶": &emojiMeta{
		Keywords: []string{"boat", "paddle", "water", "ship"},
		Name:     "canoe",
		Category: "travel_and_places",
	},
	"⚓": &emojiMeta{
		Keywords: []string{"ship", "ferry", "sea", "boat"},
		Name:     "anchor",
		Category: "travel_and_places",
	},
	"🚧": &emojiMeta{
		Keywords: []string{"wip", "progress", "caution", "warning"},
		Name:     "construction",
		Category: "travel_and_places",
	},
	"⛽": &emojiMeta{
		Keywords: []string{"gas station", "petroleum"},
		Name:     "fuelpump",
		Category: "travel_and_places",
	},
	"🚏": &emojiMeta{
		Keywords: []string{"transportation", "wait"},
		Name:     "busstop",
		Category: "travel_and_places",
	},
	"🚦": &emojiMeta{
		Keywords: []string{"transportation", "driving"},
		Name:     "vertical_traffic_light",
		Category: "travel_and_places",
	},
	"🚥": &emojiMeta{
		Keywords: []string{"transportation", "signal"},
		Name:     "traffic_light",
		Category: "travel_and_places",
	},
	"🏁": &emojiMeta{
		Keywords: []string{"contest", "finishline", "race", "gokart"},
		Name:     "checkered_flag",
		Category: "travel_and_places",
	},
	"🚢": &emojiMeta{
		Keywords: []string{"transportation", "titanic", "deploy"},
		Name:     "ship",
		Category: "travel_and_places",
	},
	"🎡": &emojiMeta{
		Keywords: []string{"photo", "carnival", "londoneye"},
		Name:     "ferris_wheel",
		Category: "travel_and_places",
	},
	"🎢": &emojiMeta{
		Keywords: []string{"carnival", "playground", "photo", "fun"},
		Name:     "roller_coaster",
		Category: "travel_and_places",
	},
	"🎠": &emojiMeta{
		Keywords: []string{"photo", "carnival"},
		Name:     "carousel_horse",
		Category: "travel_and_places",
	},
	"🏗": &emojiMeta{
		Keywords: []string{"wip", "working", "progress"},
		Name:     "building_construction",
		Category: "travel_and_places",
	},
	"🌁": &emojiMeta{
		Keywords: []string{"photo", "mountain"},
		Name:     "foggy",
		Category: "travel_and_places",
	},
	"🗼": &emojiMeta{
		Keywords: []string{"photo", "japanese"},
		Name:     "tokyo_tower",
		Category: "travel_and_places",
	},
	"🏭": &emojiMeta{
		Keywords: []string{"building", "industry", "pollution", "smoke"},
		Name:     "factory",
		Category: "travel_and_places",
	},
	"⛲": &emojiMeta{
		Keywords: []string{"photo", "summer", "water", "fresh"},
		Name:     "fountain",
		Category: "travel_and_places",
	},
	"🎑": &emojiMeta{
		Keywords: []string{"photo", "japan", "asia", "tsukimi"},
		Name:     "rice_scene",
		Category: "travel_and_places",
	},
	"⛰": &emojiMeta{
		Keywords: []string{"photo", "nature", "environment"},
		Name:     "mountain",
		Category: "travel_and_places",
	},
	"🏔": &emojiMeta{
		Keywords: []string{"photo", "nature", "environment", "winter", "cold"},
		Name:     "mountain_snow",
		Category: "travel_and_places",
	},
	"🗻": &emojiMeta{
		Keywords: []string{"photo", "mountain", "nature", "japanese"},
		Name:     "mount_fuji",
		Category: "travel_and_places",
	},
	"🌋": &emojiMeta{
		Keywords: []string{"photo", "nature", "disaster"},
		Name:     "volcano",
		Category: "travel_and_places",
	},
	"🗾": &emojiMeta{
		Keywords: []string{"nation", "country", "japanese", "asia"},
		Name:     "japan",
		Category: "travel_and_places",
	},
	"🏕": &emojiMeta{
		Keywords: []string{"photo", "outdoors", "tent"},
		Name:     "camping",
		Category: "travel_and_places",
	},
	"⛺": &emojiMeta{
		Keywords: []string{"photo", "camping", "outdoors"},
		Name:     "tent",
		Category: "travel_and_places",
	},
	"🏞": &emojiMeta{
		Keywords: []string{"photo", "environment", "nature"},
		Name:     "national_park",
		Category: "travel_and_places",
	},
	"🛣": &emojiMeta{
		Keywords: []string{"road", "cupertino", "interstate", "highway"},
		Name:     "motorway",
		Category: "travel_and_places",
	},
	"🛤": &emojiMeta{
		Keywords: []string{"train", "transportation"},
		Name:     "railway_track",
		Category: "travel_and_places",
	},
	"🌅": &emojiMeta{
		Keywords: []string{"morning", "view", "vacation", "photo"},
		Name:     "sunrise",
		Category: "travel_and_places",
	},
	"🌄": &emojiMeta{
		Keywords: []string{"view", "vacation", "photo"},
		Name:     "sunrise_over_mountains",
		Category: "travel_and_places",
	},
	"🏜": &emojiMeta{
		Keywords: []string{"photo", "warm", "saharah"},
		Name:     "desert",
		Category: "travel_and_places",
	},
	"🏖": &emojiMeta{
		Keywords: []string{"weather", "summer", "sunny", "sand", "mojito"},
		Name:     "beach_umbrella",
		Category: "travel_and_places",
	},
	"🏝": &emojiMeta{
		Keywords: []string{"photo", "tropical", "mojito"},
		Name:     "desert_island",
		Category: "travel_and_places",
	},
	"🌇": &emojiMeta{
		Keywords: []string{"photo", "good morning", "dawn"},
		Name:     "city_sunrise",
		Category: "travel_and_places",
	},
	"🌆": &emojiMeta{
		Keywords: []string{"photo", "evening", "sky", "buildings"},
		Name:     "city_sunset",
		Category: "travel_and_places",
	},
	"🏙": &emojiMeta{
		Keywords: []string{"photo", "night life", "urban"},
		Name:     "cityscape",
		Category: "travel_and_places",
	},
	"🌃": &emojiMeta{
		Keywords: []string{"evening", "city", "downtown"},
		Name:     "night_with_stars",
		Category: "travel_and_places",
	},
	"🌉": &emojiMeta{
		Keywords: []string{"photo", "sanfrancisco"},
		Name:     "bridge_at_night",
		Category: "travel_and_places",
	},
	"🌌": &emojiMeta{
		Keywords: []string{"photo", "space", "stars"},
		Name:     "milky_way",
		Category: "travel_and_places",
	},
	"🌠": &emojiMeta{
		Keywords: []string{"night", "photo"},
		Name:     "stars",
		Category: "travel_and_places",
	},
	"🎇": &emojiMeta{
		Keywords: []string{"stars", "night", "shine"},
		Name:     "sparkler",
		Category: "travel_and_places",
	},
	"🎆": &emojiMeta{
		Keywords: []string{"photo", "festival", "carnival", "congratulations"},
		Name:     "fireworks",
		Category: "travel_and_places",
	},
	"🌈": &emojiMeta{
		Keywords: []string{"nature", "happy", "unicorn_face", "photo", "sky", "spring"},
		Name:     "rainbow",
		Category: "travel_and_places",
	},
	"🏘": &emojiMeta{
		Keywords: []string{"buildings", "photo"},
		Name:     "houses",
		Category: "travel_and_places",
	},
	"🏰": &emojiMeta{
		Keywords: []string{"building", "royalty", "history"},
		Name:     "european_castle",
		Category: "travel_and_places",
	},
	"🏯": &emojiMeta{
		Keywords: []string{"photo", "building"},
		Name:     "japanese_castle",
		Category: "travel_and_places",
	},
	"🏟": &emojiMeta{
		Keywords: []string{"photo", "place", "sports", "concert", "venue"},
		Name:     "stadium",
		Category: "travel_and_places",
	},
	"🗽": &emojiMeta{
		Keywords: []string{"american", "newyork"},
		Name:     "statue_of_liberty",
		Category: "travel_and_places",
	},
	"🏠": &emojiMeta{
		Keywords: []string{"building", "home"},
		Name:     "house",
		Category: "travel_and_places",
	},
	"🏡": &emojiMeta{
		Keywords: []string{"home", "plant", "nature"},
		Name:     "house_with_garden",
		Category: "travel_and_places",
	},
	"🏚": &emojiMeta{
		Keywords: []string{"abandon", "evict", "broken", "building"},
		Name:     "derelict_house",
		Category: "travel_and_places",
	},
	"🏢": &emojiMeta{
		Keywords: []string{"building", "bureau", "work"},
		Name:     "office",
		Category: "travel_and_places",
	},
	"🏬": &emojiMeta{
		Keywords: []string{"building", "shopping", "mall"},
		Name:     "department_store",
		Category: "travel_and_places",
	},
	"🏣": &emojiMeta{
		Keywords: []string{"building", "envelope", "communication"},
		Name:     "post_office",
		Category: "travel_and_places",
	},
	"🏤": &emojiMeta{
		Keywords: []string{"building", "email"},
		Name:     "european_post_office",
		Category: "travel_and_places",
	},
	"🏥": &emojiMeta{
		Keywords: []string{"building", "health", "surgery", "doctor"},
		Name:     "hospital",
		Category: "travel_and_places",
	},
	"🏦": &emojiMeta{
		Keywords: []string{"building", "money", "sales", "cash", "business", "enterprise"},
		Name:     "bank",
		Category: "travel_and_places",
	},
	"🏨": &emojiMeta{
		Keywords: []string{"building", "accomodation", "checkin"},
		Name:     "hotel",
		Category: "travel_and_places",
	},
	"🏪": &emojiMeta{
		Keywords: []string{"building", "shopping", "groceries"},
		Name:     "convenience_store",
		Category: "travel_and_places",
	},
	"🏫": &emojiMeta{
		Keywords: []string{"building", "student", "education", "learn", "teach"},
		Name:     "school",
		Category: "travel_and_places",
	},
	"🏩": &emojiMeta{
		Keywords: []string{"like", "affection", "dating"},
		Name:     "love_hotel",
		Category: "travel_and_places",
	},
	"💒": &emojiMeta{
		Keywords: []string{"love", "like", "affection", "couple", "marriage", "bride", "groom"},
		Name:     "wedding",
		Category: "travel_and_places",
	},
	"🏛": &emojiMeta{
		Keywords: []string{"art", "culture", "history"},
		Name:     "classical_building",
		Category: "travel_and_places",
	},
	"⛪": &emojiMeta{
		Keywords: []string{"building", "religion", "christ"},
		Name:     "church",
		Category: "travel_and_places",
	},
	"🕌": &emojiMeta{
		Keywords: []string{"islam", "worship", "minaret"},
		Name:     "mosque",
		Category: "travel_and_places",
	},
	"🕍": &emojiMeta{
		Keywords: []string{"judaism", "worship", "temple", "jewish"},
		Name:     "synagogue",
		Category: "travel_and_places",
	},
	"🕋": &emojiMeta{
		Keywords: []string{"mecca", "mosque", "islam"},
		Name:     "kaaba",
		Category: "travel_and_places",
	},
	"⛩": &emojiMeta{
		Keywords: []string{"temple", "japan", "kyoto"},
		Name:     "shinto_shrine",
		Category: "travel_and_places",
	},
	"⌚": &emojiMeta{
		Keywords: []string{"time", "accessories"},
		Name:     "watch",
		Category: "objects",
	},
	"📱": &emojiMeta{
		Keywords: []string{"technology", "apple", "gadgets", "dial"},
		Name:     "iphone",
		Category: "objects",
	},
	"📲": &emojiMeta{
		Keywords: []string{"iphone", "incoming"},
		Name:     "calling",
		Category: "objects",
	},
	"💻": &emojiMeta{
		Keywords: []string{"technology", "laptop", "screen", "display", "monitor"},
		Name:     "computer",
		Category: "objects",
	},
	"⌨": &emojiMeta{
		Keywords: []string{"technology", "computer", "type", "input", "text"},
		Name:     "keyboard",
		Category: "objects",
	},
	"🖥": &emojiMeta{
		Keywords: []string{"technology", "computing", "screen"},
		Name:     "desktop_computer",
		Category: "objects",
	},
	"🖨": &emojiMeta{
		Keywords: []string{"paper", "ink"},
		Name:     "printer",
		Category: "objects",
	},
	"🖱": &emojiMeta{
		Keywords: []string{"click"},
		Name:     "computer_mouse",
		Category: "objects",
	},
	"🖲": &emojiMeta{
		Keywords: []string{"technology", "trackpad"},
		Name:     "trackball",
		Category: "objects",
	},
	"🕹": &emojiMeta{
		Keywords: []string{"game", "play"},
		Name:     "joystick",
		Category: "objects",
	},
	"🗜": &emojiMeta{
		Keywords: []string{"tool"},
		Name:     "clamp",
		Category: "objects",
	},
	"💽": &emojiMeta{
		Keywords: []string{"technology", "record", "data", "disk", "90s"},
		Name:     "minidisc",
		Category: "objects",
	},
	"💾": &emojiMeta{
		Keywords: []string{"oldschool", "technology", "save", "90s", "80s"},
		Name:     "floppy_disk",
		Category: "objects",
	},
	"💿": &emojiMeta{
		Keywords: []string{"technology", "dvd", "disk", "disc", "90s"},
		Name:     "cd",
		Category: "objects",
	},
	"📀": &emojiMeta{
		Keywords: []string{"cd", "disk", "disc"},
		Name:     "dvd",
		Category: "objects",
	},
	"📼": &emojiMeta{
		Keywords: []string{"record", "video", "oldschool", "90s", "80s"},
		Name:     "vhs",
		Category: "objects",
	},
	"📷": &emojiMeta{
		Keywords: []string{"gadgets", "photography"},
		Name:     "camera",
		Category: "objects",
	},
	"📸": &emojiMeta{
		Keywords: []string{"photography", "gadgets"},
		Name:     "camera_flash",
		Category: "objects",
	},
	"📹": &emojiMeta{
		Keywords: []string{"film", "record"},
		Name:     "video_camera",
		Category: "objects",
	},
	"🎥": &emojiMeta{
		Keywords: []string{"film", "record"},
		Name:     "movie_camera",
		Category: "objects",
	},
	"📽": &emojiMeta{
		Keywords: []string{"video", "tape", "record", "movie"},
		Name:     "film_projector",
		Category: "objects",
	},
	"🎞": &emojiMeta{
		Keywords: []string{"movie"},
		Name:     "film_strip",
		Category: "objects",
	},
	"📞": &emojiMeta{
		Keywords: []string{"technology", "communication", "dial"},
		Name:     "telephone_receiver",
		Category: "objects",
	},
	"☎️": &emojiMeta{
		Keywords: []string{"technology", "communication", "dial", "telephone"},
		Name:     "phone",
		Category: "objects",
	},
	"📟": &emojiMeta{
		Keywords: []string{"bbcall", "oldschool", "90s"},
		Name:     "pager",
		Category: "objects",
	},
	"📠": &emojiMeta{
		Keywords: []string{"communication", "technology"},
		Name:     "fax",
		Category: "objects",
	},
	"📺": &emojiMeta{
		Keywords: []string{"technology", "program", "oldschool", "show", "television"},
		Name:     "tv",
		Category: "objects",
	},
	"📻": &emojiMeta{
		Keywords: []string{"communication", "music", "podcast", "program"},
		Name:     "radio",
		Category: "objects",
	},
	"🎙": &emojiMeta{
		Keywords: []string{"sing", "recording", "artist", "talkshow"},
		Name:     "studio_microphone",
		Category: "objects",
	},
	"🎚": &emojiMeta{
		Keywords: []string{"scale"},
		Name:     "level_slider",
		Category: "objects",
	},
	"🎛": &emojiMeta{
		Keywords: []string{"dial"},
		Name:     "control_knobs",
		Category: "objects",
	},
	"🧭": &emojiMeta{
		Keywords: []string{"magnetic", "navigation", "orienteering"},
		Name:     "compass",
		Category: "objects",
	},
	"⏱": &emojiMeta{
		Keywords: []string{"time", "deadline"},
		Name:     "stopwatch",
		Category: "objects",
	},
	"⏲": &emojiMeta{
		Keywords: []string{"alarm"},
		Name:     "timer_clock",
		Category: "objects",
	},
	"⏰": &emojiMeta{
		Keywords: []string{"time", "wake"},
		Name:     "alarm_clock",
		Category: "objects",
	},
	"🕰": &emojiMeta{
		Keywords: []string{"time"},
		Name:     "mantelpiece_clock",
		Category: "objects",
	},
	"⏳": &emojiMeta{
		Keywords: []string{"oldschool", "time", "countdown"},
		Name:     "hourglass_flowing_sand",
		Category: "objects",
	},
	"⌛": &emojiMeta{
		Keywords: []string{"time", "clock", "oldschool", "limit", "exam", "quiz", "test"},
		Name:     "hourglass",
		Category: "objects",
	},
	"📡": &emojiMeta{
		Keywords: []string{"communication", "future", "radio", "space"},
		Name:     "satellite",
		Category: "objects",
	},
	"🔋": &emojiMeta{
		Keywords: []string{"power", "energy", "sustain"},
		Name:     "battery",
		Category: "objects",
	},
	"🔌": &emojiMeta{
		Keywords: []string{"charger", "power"},
		Name:     "electric_plug",
		Category: "objects",
	},
	"💡": &emojiMeta{
		Keywords: []string{"light", "electricity", "idea"},
		Name:     "bulb",
		Category: "objects",
	},
	"🔦": &emojiMeta{
		Keywords: []string{"dark", "camping", "sight", "night"},
		Name:     "flashlight",
		Category: "objects",
	},
	"🕯": &emojiMeta{
		Keywords: []string{"fire", "wax"},
		Name:     "candle",
		Category: "objects",
	},
	"🧯": &emojiMeta{
		Keywords: []string{"quench"},
		Name:     "fire_extinguisher",
		Category: "objects",
	},
	"🗑": &emojiMeta{
		Keywords: []string{"bin", "trash", "rubbish", "garbage", "toss"},
		Name:     "wastebasket",
		Category: "objects",
	},
	"🛢": &emojiMeta{
		Keywords: []string{"barrell"},
		Name:     "oil_drum",
		Category: "objects",
	},
	"💸": &emojiMeta{
		Keywords: []string{"dollar", "bills", "payment", "sale"},
		Name:     "money_with_wings",
		Category: "objects",
	},
	"💵": &emojiMeta{
		Keywords: []string{"money", "sales", "bill", "currency"},
		Name:     "dollar",
		Category: "objects",
	},
	"💴": &emojiMeta{
		Keywords: []string{"money", "sales", "japanese", "dollar", "currency"},
		Name:     "yen",
		Category: "objects",
	},
	"💶": &emojiMeta{
		Keywords: []string{"money", "sales", "dollar", "currency"},
		Name:     "euro",
		Category: "objects",
	},
	"💷": &emojiMeta{
		Keywords: []string{"british", "sterling", "money", "sales", "bills", "uk", "england", "currency"},
		Name:     "pound",
		Category: "objects",
	},
	"💰": &emojiMeta{
		Keywords: []string{"dollar", "payment", "coins", "sale"},
		Name:     "moneybag",
		Category: "objects",
	},
	"💳": &emojiMeta{
		Keywords: []string{"money", "sales", "dollar", "bill", "payment", "shopping"},
		Name:     "credit_card",
		Category: "objects",
	},
	"💎": &emojiMeta{
		Keywords: []string{"blue", "ruby", "diamond", "jewelry"},
		Name:     "gem",
		Category: "objects",
	},
	"⚖": &emojiMeta{
		Keywords: []string{"law", "fairness", "weight"},
		Name:     "balance_scale",
		Category: "objects",
	},
	"🧰": &emojiMeta{
		Keywords: []string{"tools", "diy", "fix", "maintainer", "mechanic"},
		Name:     "toolbox",
		Category: "objects",
	},
	"🔧": &emojiMeta{
		Keywords: []string{"tools", "diy", "ikea", "fix", "maintainer"},
		Name:     "wrench",
		Category: "objects",
	},
	"🔨": &emojiMeta{
		Keywords: []string{"tools", "build", "create"},
		Name:     "hammer",
		Category: "objects",
	},
	"⚒": &emojiMeta{
		Keywords: []string{"tools", "build", "create"},
		Name:     "hammer_and_pick",
		Category: "objects",
	},
	"🛠": &emojiMeta{
		Keywords: []string{"tools", "build", "create"},
		Name:     "hammer_and_wrench",
		Category: "objects",
	},
	"⛏": &emojiMeta{
		Keywords: []string{"tools", "dig"},
		Name:     "pick",
		Category: "objects",
	},
	"🔩": &emojiMeta{
		Keywords: []string{"handy", "tools", "fix"},
		Name:     "nut_and_bolt",
		Category: "objects",
	},
	"⚙": &emojiMeta{
		Keywords: []string{"cog"},
		Name:     "gear",
		Category: "objects",
	},
	"🧱": &emojiMeta{
		Keywords: []string{"bricks"},
		Name:     "brick",
		Category: "objects",
	},
	"⛓": &emojiMeta{
		Keywords: []string{"lock", "arrest"},
		Name:     "chains",
		Category: "objects",
	},
	"🧲": &emojiMeta{
		Keywords: []string{"attraction", "magnetic"},
		Name:     "magnet",
		Category: "objects",
	},
	"🔫": &emojiMeta{
		Keywords: []string{"violence", "weapon", "pistol", "revolver"},
		Name:     "gun",
		Category: "objects",
	},
	"💣": &emojiMeta{
		Keywords: []string{"boom", "explode", "explosion", "terrorism"},
		Name:     "bomb",
		Category: "objects",
	},
	"🧨": &emojiMeta{
		Keywords: []string{"dynamite", "boom", "explode", "explosion", "explosive"},
		Name:     "firecracker",
		Category: "objects",
	},
	"🔪": &emojiMeta{
		Keywords: []string{"knife", "blade", "cutlery", "kitchen", "weapon"},
		Name:     "hocho",
		Category: "objects",
	},
	"🗡": &emojiMeta{
		Keywords: []string{"weapon"},
		Name:     "dagger",
		Category: "objects",
	},
	"⚔": &emojiMeta{
		Keywords: []string{"weapon"},
		Name:     "crossed_swords",
		Category: "objects",
	},
	"🛡": &emojiMeta{
		Keywords: []string{"protection", "security"},
		Name:     "shield",
		Category: "objects",
	},
	"🚬": &emojiMeta{
		Keywords: []string{"kills", "tobacco", "cigarette", "joint", "smoke"},
		Name:     "smoking",
		Category: "objects",
	},
	"☠": &emojiMeta{
		Keywords: []string{"poison", "danger", "deadly", "scary", "death", "pirate", "evil"},
		Name:     "skull_and_crossbones",
		Category: "objects",
	},
	"⚰": &emojiMeta{
		Keywords: []string{"vampire", "dead", "die", "death", "rip", "graveyard", "cemetery", "casket", "funeral", "box"},
		Name:     "coffin",
		Category: "objects",
	},
	"⚱": &emojiMeta{
		Keywords: []string{"dead", "die", "death", "rip", "ashes"},
		Name:     "funeral_urn",
		Category: "objects",
	},
	"🏺": &emojiMeta{
		Keywords: []string{"vase", "jar"},
		Name:     "amphora",
		Category: "objects",
	},
	"🔮": &emojiMeta{
		Keywords: []string{"disco", "party", "magic", "circus", "fortune_teller"},
		Name:     "crystal_ball",
		Category: "objects",
	},
	"📿": &emojiMeta{
		Keywords: []string{"dhikr", "religious"},
		Name:     "prayer_beads",
		Category: "objects",
	},
	"🧿": &emojiMeta{
		Keywords: []string{"bead", "charm"},
		Name:     "nazar_amulet",
		Category: "objects",
	},
	"💈": &emojiMeta{
		Keywords: []string{"hair", "salon", "style"},
		Name:     "barber",
		Category: "objects",
	},
	"⚗": &emojiMeta{
		Keywords: []string{"distilling", "science", "experiment", "chemistry"},
		Name:     "alembic",
		Category: "objects",
	},
	"🔭": &emojiMeta{
		Keywords: []string{"stars", "space", "zoom", "science", "astronomy"},
		Name:     "telescope",
		Category: "objects",
	},
	"🔬": &emojiMeta{
		Keywords: []string{"laboratory", "experiment", "zoomin", "science", "study"},
		Name:     "microscope",
		Category: "objects",
	},
	"🕳": &emojiMeta{
		Keywords: []string{"embarrassing"},
		Name:     "hole",
		Category: "objects",
	},
	"💊": &emojiMeta{
		Keywords: []string{"health", "medicine", "doctor", "pharmacy", "drug"},
		Name:     "pill",
		Category: "objects",
	},
	"💉": &emojiMeta{
		Keywords: []string{"health", "hospital", "drugs", "blood", "medicine", "needle", "doctor", "nurse"},
		Name:     "syringe",
		Category: "objects",
	},
	"🧬": &emojiMeta{
		Keywords: []string{"biologist", "genetics", "life"},
		Name:     "dna",
		Category: "objects",
	},
	"🦠": &emojiMeta{
		Keywords: []string{"amoeba", "bacteria", "germs"},
		Name:     "microbe",
		Category: "objects",
	},
	"🧫": &emojiMeta{
		Keywords: []string{"bacteria", "biology", "culture", "lab"},
		Name:     "petri_dish",
		Category: "objects",
	},
	"🧪": &emojiMeta{
		Keywords: []string{"chemistry", "experiment", "lab", "science"},
		Name:     "test_tube",
		Category: "objects",
	},
	"🌡": &emojiMeta{
		Keywords: []string{"weather", "temperature", "hot", "cold"},
		Name:     "thermometer",
		Category: "objects",
	},
	"🧹": &emojiMeta{
		Keywords: []string{"cleaning", "sweeping", "witch"},
		Name:     "broom",
		Category: "objects",
	},
	"🧺": &emojiMeta{
		Keywords: []string{"laundry"},
		Name:     "basket",
		Category: "objects",
	},
	"🧻": &emojiMeta{
		Keywords: []string{"roll"},
		Name:     "toilet_paper",
		Category: "objects",
	},
	"🏷": &emojiMeta{
		Keywords: []string{"sale", "tag"},
		Name:     "label",
		Category: "objects",
	},
	"🔖": &emojiMeta{
		Keywords: []string{"favorite", "label", "save"},
		Name:     "bookmark",
		Category: "objects",
	},
	"🚽": &emojiMeta{
		Keywords: []string{"restroom", "wc", "washroom", "bathroom", "potty"},
		Name:     "toilet",
		Category: "objects",
	},
	"🚿": &emojiMeta{
		Keywords: []string{"clean", "water", "bathroom"},
		Name:     "shower",
		Category: "objects",
	},
	"🛁": &emojiMeta{
		Keywords: []string{"clean", "shower", "bathroom"},
		Name:     "bathtub",
		Category: "objects",
	},
	"🧼": &emojiMeta{
		Keywords: []string{"bar", "bathing", "cleaning", "lather"},
		Name:     "soap",
		Category: "objects",
	},
	"🧽": &emojiMeta{
		Keywords: []string{"absorbing", "cleaning", "porous"},
		Name:     "sponge",
		Category: "objects",
	},
	"🧴": &emojiMeta{
		Keywords: []string{"moisturizer", "sunscreen"},
		Name:     "lotion_bottle",
		Category: "objects",
	},
	"🔑": &emojiMeta{
		Keywords: []string{"lock", "door", "password"},
		Name:     "key",
		Category: "objects",
	},
	"🗝": &emojiMeta{
		Keywords: []string{"lock", "door", "password"},
		Name:     "old_key",
		Category: "objects",
	},
	"🛋": &emojiMeta{
		Keywords: []string{"read", "chill"},
		Name:     "couch_and_lamp",
		Category: "objects",
	},
	"🛌": &emojiMeta{
		Keywords: []string{"bed", "rest"},
		Name:     "sleeping_bed",
		Category: "objects",
	},
	"🛏": &emojiMeta{
		Keywords: []string{"sleep", "rest"},
		Name:     "bed",
		Category: "objects",
	},
	"🚪": &emojiMeta{
		Keywords: []string{"house", "entry", "exit"},
		Name:     "door",
		Category: "objects",
	},
	"🛎": &emojiMeta{
		Keywords: []string{"service"},
		Name:     "bellhop_bell",
		Category: "objects",
	},
	"🧸": &emojiMeta{
		Keywords: []string{"plush", "stuffed"},
		Name:     "teddy_bear",
		Category: "objects",
	},
	"🖼": &emojiMeta{
		Keywords: []string{"photography"},
		Name:     "framed_picture",
		Category: "objects",
	},
	"🗺": &emojiMeta{
		Keywords: []string{"location", "direction"},
		Name:     "world_map",
		Category: "objects",
	},
	"⛱": &emojiMeta{
		Keywords: []string{"weather", "summer"},
		Name:     "parasol_on_ground",
		Category: "objects",
	},
	"🗿": &emojiMeta{
		Keywords: []string{"rock", "easter island", "moai"},
		Name:     "moyai",
		Category: "objects",
	},
	"🛍": &emojiMeta{
		Keywords: []string{"mall", "buy", "purchase"},
		Name:     "shopping",
		Category: "objects",
	},
	"🛒": &emojiMeta{
		Keywords: []string{"trolley"},
		Name:     "shopping_cart",
		Category: "objects",
	},
	"🎈": &emojiMeta{
		Keywords: []string{"party", "celebration", "birthday", "circus"},
		Name:     "balloon",
		Category: "objects",
	},
	"🎏": &emojiMeta{
		Keywords: []string{"fish", "japanese", "koinobori", "carp", "banner"},
		Name:     "flag",
		Category: "objects",
	},
	"🎀": &emojiMeta{
		Keywords: []string{"decoration", "pink", "girl", "bowtie"},
		Name:     "ribbon",
		Category: "objects",
	},
	"🎁": &emojiMeta{
		Keywords: []string{"present", "birthday", "christmas", "xmas"},
		Name:     "gift",
		Category: "objects",
	},
	"🎊": &emojiMeta{
		Keywords: []string{"festival", "party", "birthday", "circus"},
		Name:     "confetti_ball",
		Category: "objects",
	},
	"🎉": &emojiMeta{
		Keywords: []string{"party", "congratulations", "birthday", "magic", "circus", "celebration"},
		Name:     "tada",
		Category: "objects",
	},
	"🎎": &emojiMeta{
		Keywords: []string{"japanese", "toy", "kimono"},
		Name:     "dolls",
		Category: "objects",
	},
	"🎐": &emojiMeta{
		Keywords: []string{"nature", "ding", "spring", "bell"},
		Name:     "wind_chime",
		Category: "objects",
	},
	"🎌": &emojiMeta{
		Keywords: []string{"japanese", "nation", "country", "border"},
		Name:     "crossed_flags",
		Category: "objects",
	},
	"🏮": &emojiMeta{
		Keywords: []string{"light", "paper", "halloween", "spooky"},
		Name:     "izakaya_lantern",
		Category: "objects",
	},
	"🧧": &emojiMeta{
		Keywords: []string{"gift"},
		Name:     "red_envelope",
		Category: "objects",
	},
	"✉️": &emojiMeta{
		Keywords: []string{"letter", "postal", "inbox", "communication"},
		Name:     "email",
		Category: "objects",
	},
	"📩": &emojiMeta{
		Keywords: []string{"email", "communication"},
		Name:     "envelope_with_arrow",
		Category: "objects",
	},
	"📨": &emojiMeta{
		Keywords: []string{"email", "inbox"},
		Name:     "incoming_envelope",
		Category: "objects",
	},
	"📧": &emojiMeta{
		Keywords: []string{"communication", "inbox"},
		Name:     "e-mail",
		Category: "objects",
	},
	"💌": &emojiMeta{
		Keywords: []string{"email", "like", "affection", "envelope", "valentines"},
		Name:     "love_letter",
		Category: "objects",
	},
	"📮": &emojiMeta{
		Keywords: []string{"email", "letter", "envelope"},
		Name:     "postbox",
		Category: "objects",
	},
	"📪": &emojiMeta{
		Keywords: []string{"email", "communication", "inbox"},
		Name:     "mailbox_closed",
		Category: "objects",
	},
	"📫": &emojiMeta{
		Keywords: []string{"email", "inbox", "communication"},
		Name:     "mailbox",
		Category: "objects",
	},
	"📬": &emojiMeta{
		Keywords: []string{"email", "inbox", "communication"},
		Name:     "mailbox_with_mail",
		Category: "objects",
	},
	"📭": &emojiMeta{
		Keywords: []string{"email", "inbox"},
		Name:     "mailbox_with_no_mail",
		Category: "objects",
	},
	"📦": &emojiMeta{
		Keywords: []string{"mail", "gift", "cardboard", "box", "moving"},
		Name:     "package",
		Category: "objects",
	},
	"📯": &emojiMeta{
		Keywords: []string{"instrument", "music"},
		Name:     "postal_horn",
		Category: "objects",
	},
	"📥": &emojiMeta{
		Keywords: []string{"email", "documents"},
		Name:     "inbox_tray",
		Category: "objects",
	},
	"📤": &emojiMeta{
		Keywords: []string{"inbox", "email"},
		Name:     "outbox_tray",
		Category: "objects",
	},
	"📜": &emojiMeta{
		Keywords: []string{"documents", "ancient", "history", "paper"},
		Name:     "scroll",
		Category: "objects",
	},
	"📃": &emojiMeta{
		Keywords: []string{"documents", "office", "paper"},
		Name:     "page_with_curl",
		Category: "objects",
	},
	"📑": &emojiMeta{
		Keywords: []string{"favorite", "save", "order", "tidy"},
		Name:     "bookmark_tabs",
		Category: "objects",
	},
	"🧾": &emojiMeta{
		Keywords: []string{"accounting", "expenses"},
		Name:     "receipt",
		Category: "objects",
	},
	"📊": &emojiMeta{
		Keywords: []string{"graph", "presentation", "stats"},
		Name:     "bar_chart",
		Category: "objects",
	},
	"📈": &emojiMeta{
		Keywords: []string{"graph", "presentation", "stats", "recovery", "business", "economics", "money", "sales", "good", "success"},
		Name:     "chart_with_upwards_trend",
		Category: "objects",
	},
	"📉": &emojiMeta{
		Keywords: []string{"graph", "presentation", "stats", "recession", "business", "economics", "money", "sales", "bad", "failure"},
		Name:     "chart_with_downwards_trend",
		Category: "objects",
	},
	"📄": &emojiMeta{
		Keywords: []string{"documents", "office", "paper", "information"},
		Name:     "page_facing_up",
		Category: "objects",
	},
	"📅": &emojiMeta{
		Keywords: []string{"calendar", "schedule"},
		Name:     "date",
		Category: "objects",
	},
	"📆": &emojiMeta{
		Keywords: []string{"schedule", "date", "planning"},
		Name:     "calendar",
		Category: "objects",
	},
	"🗓": &emojiMeta{
		Keywords: []string{"date", "schedule", "planning"},
		Name:     "spiral_calendar",
		Category: "objects",
	},
	"📇": &emojiMeta{
		Keywords: []string{"business", "stationery"},
		Name:     "card_index",
		Category: "objects",
	},
	"🗃": &emojiMeta{
		Keywords: []string{"business", "stationery"},
		Name:     "card_file_box",
		Category: "objects",
	},
	"🗳": &emojiMeta{
		Keywords: []string{"election", "vote"},
		Name:     "ballot_box",
		Category: "objects",
	},
	"🗄": &emojiMeta{
		Keywords: []string{"filing", "organizing"},
		Name:     "file_cabinet",
		Category: "objects",
	},
	"📋": &emojiMeta{
		Keywords: []string{"stationery", "documents"},
		Name:     "clipboard",
		Category: "objects",
	},
	"🗒": &emojiMeta{
		Keywords: []string{"memo", "stationery"},
		Name:     "spiral_notepad",
		Category: "objects",
	},
	"📁": &emojiMeta{
		Keywords: []string{"documents", "business", "office"},
		Name:     "file_folder",
		Category: "objects",
	},
	"📂": &emojiMeta{
		Keywords: []string{"documents", "load"},
		Name:     "open_file_folder",
		Category: "objects",
	},
	"🗂": &emojiMeta{
		Keywords: []string{"organizing", "business", "stationery"},
		Name:     "card_index_dividers",
		Category: "objects",
	},
	"🗞": &emojiMeta{
		Keywords: []string{"press", "headline"},
		Name:     "newspaper_roll",
		Category: "objects",
	},
	"📰": &emojiMeta{
		Keywords: []string{"press", "headline"},
		Name:     "newspaper",
		Category: "objects",
	},
	"📓": &emojiMeta{
		Keywords: []string{"stationery", "record", "notes", "paper", "study"},
		Name:     "notebook",
		Category: "objects",
	},
	"📕": &emojiMeta{
		Keywords: []string{"read", "library", "knowledge", "textbook", "learn"},
		Name:     "closed_book",
		Category: "objects",
	},
	"📗": &emojiMeta{
		Keywords: []string{"read", "library", "knowledge", "study"},
		Name:     "green_book",
		Category: "objects",
	},
	"📘": &emojiMeta{
		Keywords: []string{"read", "library", "knowledge", "learn", "study"},
		Name:     "blue_book",
		Category: "objects",
	},
	"📙": &emojiMeta{
		Keywords: []string{"read", "library", "knowledge", "textbook", "study"},
		Name:     "orange_book",
		Category: "objects",
	},
	"📔": &emojiMeta{
		Keywords: []string{"classroom", "notes", "record", "paper", "study"},
		Name:     "notebook_with_decorative_cover",
		Category: "objects",
	},
	"📒": &emojiMeta{
		Keywords: []string{"notes", "paper"},
		Name:     "ledger",
		Category: "objects",
	},
	"📚": &emojiMeta{
		Keywords: []string{"literature", "library", "study"},
		Name:     "books",
		Category: "objects",
	},
	"📖": &emojiMeta{
		Keywords: []string{"book", "read", "library", "knowledge", "literature", "learn", "study"},
		Name:     "open_book",
		Category: "objects",
	},
	"🧷": &emojiMeta{
		Keywords: []string{"diaper"},
		Name:     "safety_pin",
		Category: "objects",
	},
	"🔗": &emojiMeta{
		Keywords: []string{"rings", "url"},
		Name:     "link",
		Category: "objects",
	},
	"📎": &emojiMeta{
		Keywords: []string{"documents", "stationery"},
		Name:     "paperclip",
		Category: "objects",
	},
	"🖇": &emojiMeta{
		Keywords: []string{"documents", "stationery"},
		Name:     "paperclips",
		Category: "objects",
	},
	"✂️": &emojiMeta{
		Keywords: []string{"stationery", "cut"},
		Name:     "scissors",
		Category: "objects",
	},
	"📐": &emojiMeta{
		Keywords: []string{"stationery", "math", "architect", "sketch"},
		Name:     "triangular_ruler",
		Category: "objects",
	},
	"📏": &emojiMeta{
		Keywords: []string{"stationery", "calculate", "length", "math", "school", "drawing", "architect", "sketch"},
		Name:     "straight_ruler",
		Category: "objects",
	},
	"🧮": &emojiMeta{
		Keywords: []string{"calculation"},
		Name:     "abacus",
		Category: "objects",
	},
	"📌": &emojiMeta{
		Keywords: []string{"stationery", "mark", "here"},
		Name:     "pushpin",
		Category: "objects",
	},
	"📍": &emojiMeta{
		Keywords: []string{"stationery", "location", "map", "here"},
		Name:     "round_pushpin",
		Category: "objects",
	},
	"🚩": &emojiMeta{
		Keywords: []string{"mark", "milestone", "place"},
		Name:     "triangular_flag_on_post",
		Category: "objects",
	},
	"🏳": &emojiMeta{
		Keywords: []string{"losing", "loser", "lost", "surrender", "give up", "fail"},
		Name:     "white_flag",
		Category: "objects",
	},
	"🏴": &emojiMeta{
		Keywords: []string{"pirate"},
		Name:     "black_flag",
		Category: "objects",
	},
	"🏳️‍🌈": &emojiMeta{
		Keywords: []string{"flag", "rainbow", "pride", "gay", "lgbt", "glbt", "queer", "homosexual", "lesbian", "bisexual", "transgender"},
		Name:     "rainbow_flag",
		Category: "objects",
	},
	"🔐": &emojiMeta{
		Keywords: []string{"security", "privacy"},
		Name:     "closed_lock_with_key",
		Category: "objects",
	},
	"🔒": &emojiMeta{
		Keywords: []string{"security", "password", "padlock"},
		Name:     "lock",
		Category: "objects",
	},
	"🔓": &emojiMeta{
		Keywords: []string{"privacy", "security"},
		Name:     "unlock",
		Category: "objects",
	},
	"🔏": &emojiMeta{
		Keywords: []string{"security", "secret"},
		Name:     "lock_with_ink_pen",
		Category: "objects",
	},
	"🖊": &emojiMeta{
		Keywords: []string{"stationery", "writing", "write"},
		Name:     "pen",
		Category: "objects",
	},
	"🖋": &emojiMeta{
		Keywords: []string{"stationery", "writing", "write"},
		Name:     "fountain_pen",
		Category: "objects",
	},
	"✒️": &emojiMeta{
		Keywords: []string{"pen", "stationery", "writing", "write"},
		Name:     "black_nib",
		Category: "objects",
	},
	"📝": &emojiMeta{
		Keywords: []string{"write", "documents", "stationery", "pencil", "paper", "writing", "legal", "exam", "quiz", "test", "study", "compose"},
		Name:     "memo",
		Category: "objects",
	},
	"✏️": &emojiMeta{
		Keywords: []string{"stationery", "write", "paper", "writing", "school", "study"},
		Name:     "pencil2",
		Category: "objects",
	},
	"🖍": &emojiMeta{
		Keywords: []string{"drawing", "creativity"},
		Name:     "crayon",
		Category: "objects",
	},
	"🖌": &emojiMeta{
		Keywords: []string{"drawing", "creativity", "art"},
		Name:     "paintbrush",
		Category: "objects",
	},
	"🔍": &emojiMeta{
		Keywords: []string{"search", "zoom", "find", "detective"},
		Name:     "mag",
		Category: "objects",
	},
	"🔎": &emojiMeta{
		Keywords: []string{"search", "zoom", "find", "detective"},
		Name:     "mag_right",
		Category: "objects",
	},
	"❤️": &emojiMeta{
		Keywords: []string{"love", "like", "valentines"},
		Name:     "heart",
		Category: "symbols",
	},
	"🧡": &emojiMeta{
		Keywords: []string{"love", "like", "affection", "valentines"},
		Name:     "orange_heart",
		Category: "symbols",
	},
	"💛": &emojiMeta{
		Keywords: []string{"love", "like", "affection", "valentines"},
		Name:     "yellow_heart",
		Category: "symbols",
	},
	"💚": &emojiMeta{
		Keywords: []string{"love", "like", "affection", "valentines"},
		Name:     "green_heart",
		Category: "symbols",
	},
	"💙": &emojiMeta{
		Keywords: []string{"love", "like", "affection", "valentines"},
		Name:     "blue_heart",
		Category: "symbols",
	},
	"💜": &emojiMeta{
		Keywords: []string{"love", "like", "affection", "valentines"},
		Name:     "purple_heart",
		Category: "symbols",
	},
	"🖤": &emojiMeta{
		Keywords: []string{"evil"},
		Name:     "black_heart",
		Category: "symbols",
	},
	"💔": &emojiMeta{
		Keywords: []string{"sad", "sorry", "break", "heart", "heartbreak"},
		Name:     "broken_heart",
		Category: "symbols",
	},
	"❣": &emojiMeta{
		Keywords: []string{"decoration", "love"},
		Name:     "heavy_heart_exclamation",
		Category: "symbols",
	},
	"💕": &emojiMeta{
		Keywords: []string{"love", "like", "affection", "valentines", "heart"},
		Name:     "two_hearts",
		Category: "symbols",
	},
	"💞": &emojiMeta{
		Keywords: []string{"love", "like", "affection", "valentines"},
		Name:     "revolving_hearts",
		Category: "symbols",
	},
	"💓": &emojiMeta{
		Keywords: []string{"love", "like", "affection", "valentines", "pink", "heart"},
		Name:     "heartbeat",
		Category: "symbols",
	},
	"💗": &emojiMeta{
		Keywords: []string{"like", "love", "affection", "valentines", "pink"},
		Name:     "heartpulse",
		Category: "symbols",
	},
	"💖": &emojiMeta{
		Keywords: []string{"love", "like", "affection", "valentines"},
		Name:     "sparkling_heart",
		Category: "symbols",
	},
	"💘": &emojiMeta{
		Keywords: []string{"love", "like", "heart", "affection", "valentines"},
		Name:     "cupid",
		Category: "symbols",
	},
	"💝": &emojiMeta{
		Keywords: []string{"love", "valentines"},
		Name:     "gift_heart",
		Category: "symbols",
	},
	"💟": &emojiMeta{
		Keywords: []string{"purple-square", "love", "like"},
		Name:     "heart_decoration",
		Category: "symbols",
	},
	"☮": &emojiMeta{
		Keywords: []string{"hippie"},
		Name:     "peace_symbol",
		Category: "symbols",
	},
	"✝": &emojiMeta{
		Keywords: []string{"christianity"},
		Name:     "latin_cross",
		Category: "symbols",
	},
	"☪": &emojiMeta{
		Keywords: []string{"islam"},
		Name:     "star_and_crescent",
		Category: "symbols",
	},
	"🕉": &emojiMeta{
		Keywords: []string{"hinduism", "buddhism", "sikhism", "jainism"},
		Name:     "om",
		Category: "symbols",
	},
	"☸": &emojiMeta{
		Keywords: []string{"hinduism", "buddhism", "sikhism", "jainism"},
		Name:     "wheel_of_dharma",
		Category: "symbols",
	},
	"✡": &emojiMeta{
		Keywords: []string{"judaism"},
		Name:     "star_of_david",
		Category: "symbols",
	},
	"🔯": &emojiMeta{
		Keywords: []string{"purple-square", "religion", "jewish", "hexagram"},
		Name:     "six_pointed_star",
		Category: "symbols",
	},
	"🕎": &emojiMeta{
		Keywords: []string{"hanukkah", "candles", "jewish"},
		Name:     "menorah",
		Category: "symbols",
	},
	"☯": &emojiMeta{
		Keywords: []string{"balance"},
		Name:     "yin_yang",
		Category: "symbols",
	},
	"☦": &emojiMeta{
		Keywords: []string{"suppedaneum", "religion"},
		Name:     "orthodox_cross",
		Category: "symbols",
	},
	"🛐": &emojiMeta{
		Keywords: []string{"religion", "church", "temple", "prayer"},
		Name:     "place_of_worship",
		Category: "symbols",
	},
	"⛎": &emojiMeta{
		Keywords: []string{"sign", "purple-square", "constellation", "astrology"},
		Name:     "ophiuchus",
		Category: "symbols",
	},
	"♈": &emojiMeta{
		Keywords: []string{"sign", "purple-square", "zodiac", "astrology"},
		Name:     "aries",
		Category: "symbols",
	},
	"♉": &emojiMeta{
		Keywords: []string{"purple-square", "sign", "zodiac", "astrology"},
		Name:     "taurus",
		Category: "symbols",
	},
	"♊": &emojiMeta{
		Keywords: []string{"sign", "zodiac", "purple-square", "astrology"},
		Name:     "gemini",
		Category: "symbols",
	},
	"♋": &emojiMeta{
		Keywords: []string{"sign", "zodiac", "purple-square", "astrology"},
		Name:     "cancer",
		Category: "symbols",
	},
	"♌": &emojiMeta{
		Keywords: []string{"sign", "purple-square", "zodiac", "astrology"},
		Name:     "leo",
		Category: "symbols",
	},
	"♍": &emojiMeta{
		Keywords: []string{"sign", "zodiac", "purple-square", "astrology"},
		Name:     "virgo",
		Category: "symbols",
	},
	"♎": &emojiMeta{
		Keywords: []string{"sign", "purple-square", "zodiac", "astrology"},
		Name:     "libra",
		Category: "symbols",
	},
	"♏": &emojiMeta{
		Keywords: []string{"sign", "zodiac", "purple-square", "astrology", "scorpio"},
		Name:     "scorpius",
		Category: "symbols",
	},
	"♐": &emojiMeta{
		Keywords: []string{"sign", "zodiac", "purple-square", "astrology"},
		Name:     "sagittarius",
		Category: "symbols",
	},
	"♑": &emojiMeta{
		Keywords: []string{"sign", "zodiac", "purple-square", "astrology"},
		Name:     "capricorn",
		Category: "symbols",
	},
	"♒": &emojiMeta{
		Keywords: []string{"sign", "purple-square", "zodiac", "astrology"},
		Name:     "aquarius",
		Category: "symbols",
	},
	"♓": &emojiMeta{
		Keywords: []string{"purple-square", "sign", "zodiac", "astrology"},
		Name:     "pisces",
		Category: "symbols",
	},
	"🆔": &emojiMeta{
		Keywords: []string{"purple-square", "words"},
		Name:     "id",
		Category: "symbols",
	},
	"⚛": &emojiMeta{
		Keywords: []string{"science", "physics", "chemistry"},
		Name:     "atom_symbol",
		Category: "symbols",
	},
	"🈳": &emojiMeta{
		Keywords: []string{"kanji", "japanese", "chinese", "empty", "sky", "blue-square"},
		Name:     "u7a7a",
		Category: "symbols",
	},
	"🈹": &emojiMeta{
		Keywords: []string{"cut", "divide", "chinese", "kanji", "pink-square"},
		Name:     "u5272",
		Category: "symbols",
	},
	"☢": &emojiMeta{
		Keywords: []string{"nuclear", "danger"},
		Name:     "radioactive",
		Category: "symbols",
	},
	"☣": &emojiMeta{
		Keywords: []string{"danger"},
		Name:     "biohazard",
		Category: "symbols",
	},
	"📴": &emojiMeta{
		Keywords: []string{"mute", "orange-square", "silence", "quiet"},
		Name:     "mobile_phone_off",
		Category: "symbols",
	},
	"📳": &emojiMeta{
		Keywords: []string{"orange-square", "phone"},
		Name:     "vibration_mode",
		Category: "symbols",
	},
	"🈶": &emojiMeta{
		Keywords: []string{"orange-square", "chinese", "have", "kanji"},
		Name:     "u6709",
		Category: "symbols",
	},
	"🈚": &emojiMeta{
		Keywords: []string{"nothing", "chinese", "kanji", "japanese", "orange-square"},
		Name:     "u7121",
		Category: "symbols",
	},
	"🈸": &emojiMeta{
		Keywords: []string{"chinese", "japanese", "kanji", "orange-square"},
		Name:     "u7533",
		Category: "symbols",
	},
	"🈺": &emojiMeta{
		Keywords: []string{"japanese", "opening hours", "orange-square"},
		Name:     "u55b6",
		Category: "symbols",
	},
	"🈷️": &emojiMeta{
		Keywords: []string{"chinese", "month", "moon", "japanese", "orange-square", "kanji"},
		Name:     "u6708",
		Category: "symbols",
	},
	"✴️": &emojiMeta{
		Keywords: []string{"orange-square", "shape", "polygon"},
		Name:     "eight_pointed_black_star",
		Category: "symbols",
	},
	"🆚": &emojiMeta{
		Keywords: []string{"words", "orange-square"},
		Name:     "vs",
		Category: "symbols",
	},
	"🉑": &emojiMeta{
		Keywords: []string{"ok", "good", "chinese", "kanji", "agree", "yes", "orange-circle"},
		Name:     "accept",
		Category: "symbols",
	},
	"💮": &emojiMeta{
		Keywords: []string{"japanese", "spring"},
		Name:     "white_flower",
		Category: "symbols",
	},
	"🉐": &emojiMeta{
		Keywords: []string{"chinese", "kanji", "obtain", "get", "circle"},
		Name:     "ideograph_advantage",
		Category: "symbols",
	},
	"㊙️": &emojiMeta{
		Keywords: []string{"privacy", "chinese", "sshh", "kanji", "red-circle"},
		Name:     "secret",
		Category: "symbols",
	},
	"㊗️": &emojiMeta{
		Keywords: []string{"chinese", "kanji", "japanese", "red-circle"},
		Name:     "congratulations",
		Category: "symbols",
	},
	"🈴": &emojiMeta{
		Keywords: []string{"japanese", "chinese", "join", "kanji", "red-square"},
		Name:     "u5408",
		Category: "symbols",
	},
	"🈵": &emojiMeta{
		Keywords: []string{"full", "chinese", "japanese", "red-square", "kanji"},
		Name:     "u6e80",
		Category: "symbols",
	},
	"🈲": &emojiMeta{
		Keywords: []string{"kanji", "japanese", "chinese", "forbidden", "limit", "restricted", "red-square"},
		Name:     "u7981",
		Category: "symbols",
	},
	"🅰️": &emojiMeta{
		Keywords: []string{"red-square", "alphabet", "letter"},
		Name:     "a",
		Category: "symbols",
	},
	"🅱️": &emojiMeta{
		Keywords: []string{"red-square", "alphabet", "letter"},
		Name:     "b",
		Category: "symbols",
	},
	"🆎": &emojiMeta{
		Keywords: []string{"red-square", "alphabet"},
		Name:     "ab",
		Category: "symbols",
	},
	"🆑": &emojiMeta{
		Keywords: []string{"alphabet", "words", "red-square"},
		Name:     "cl",
		Category: "symbols",
	},
	"🅾️": &emojiMeta{
		Keywords: []string{"alphabet", "red-square", "letter"},
		Name:     "o2",
		Category: "symbols",
	},
	"🆘": &emojiMeta{
		Keywords: []string{"help", "red-square", "words", "emergency", "911"},
		Name:     "sos",
		Category: "symbols",
	},
	"⛔": &emojiMeta{
		Keywords: []string{"limit", "security", "privacy", "bad", "denied", "stop", "circle"},
		Name:     "no_entry",
		Category: "symbols",
	},
	"📛": &emojiMeta{
		Keywords: []string{"fire", "forbid"},
		Name:     "name_badge",
		Category: "symbols",
	},
	"🚫": &emojiMeta{
		Keywords: []string{"forbid", "stop", "limit", "denied", "disallow", "circle"},
		Name:     "no_entry_sign",
		Category: "symbols",
	},
	"❌": &emojiMeta{
		Keywords: []string{"no", "delete", "remove", "cancel", "red"},
		Name:     "x",
		Category: "symbols",
	},
	"⭕": &emojiMeta{
		Keywords: []string{"circle", "round"},
		Name:     "o",
		Category: "symbols",
	},
	"🛑": &emojiMeta{
		Keywords: []string{"stop"},
		Name:     "stop_sign",
		Category: "symbols",
	},
	"💢": &emojiMeta{
		Keywords: []string{"angry", "mad"},
		Name:     "anger",
		Category: "symbols",
	},
	"♨️": &emojiMeta{
		Keywords: []string{"bath", "warm", "relax"},
		Name:     "hotsprings",
		Category: "symbols",
	},
	"🚷": &emojiMeta{
		Keywords: []string{"rules", "crossing", "walking", "circle"},
		Name:     "no_pedestrians",
		Category: "symbols",
	},
	"🚯": &emojiMeta{
		Keywords: []string{"trash", "bin", "garbage", "circle"},
		Name:     "do_not_litter",
		Category: "symbols",
	},
	"🚳": &emojiMeta{
		Keywords: []string{"cyclist", "prohibited", "circle"},
		Name:     "no_bicycles",
		Category: "symbols",
	},
	"🚱": &emojiMeta{
		Keywords: []string{"drink", "faucet", "tap", "circle"},
		Name:     "non-potable_water",
		Category: "symbols",
	},
	"🔞": &emojiMeta{
		Keywords: []string{"18", "drink", "pub", "night", "minor", "circle"},
		Name:     "underage",
		Category: "symbols",
	},
	"📵": &emojiMeta{
		Keywords: []string{"iphone", "mute", "circle"},
		Name:     "no_mobile_phones",
		Category: "symbols",
	},
	"❗": &emojiMeta{
		Keywords: []string{"heavy_exclamation_mark", "danger", "surprise", "punctuation", "wow", "warning"},
		Name:     "exclamation",
		Category: "symbols",
	},
	"❕": &emojiMeta{
		Keywords: []string{"surprise", "punctuation", "gray", "wow", "warning"},
		Name:     "grey_exclamation",
		Category: "symbols",
	},
	"❓": &emojiMeta{
		Keywords: []string{"doubt", "confused"},
		Name:     "question",
		Category: "symbols",
	},
	"❔": &emojiMeta{
		Keywords: []string{"doubts", "gray", "huh", "confused"},
		Name:     "grey_question",
		Category: "symbols",
	},
	"‼️": &emojiMeta{
		Keywords: []string{"exclamation", "surprise"},
		Name:     "bangbang",
		Category: "symbols",
	},
	"⁉️": &emojiMeta{
		Keywords: []string{"wat", "punctuation", "surprise"},
		Name:     "interrobang",
		Category: "symbols",
	},
	"💯": &emojiMeta{
		Keywords: []string{"score", "perfect", "numbers", "century", "exam", "quiz", "test", "pass", "hundred"},
		Name:     "100",
		Category: "symbols",
	},
	"🔅": &emojiMeta{
		Keywords: []string{"sun", "afternoon", "warm", "summer"},
		Name:     "low_brightness",
		Category: "symbols",
	},
	"🔆": &emojiMeta{
		Keywords: []string{"sun", "light"},
		Name:     "high_brightness",
		Category: "symbols",
	},
	"🔱": &emojiMeta{
		Keywords: []string{"weapon", "spear"},
		Name:     "trident",
		Category: "symbols",
	},
	"⚜": &emojiMeta{
		Keywords: []string{"decorative", "scout"},
		Name:     "fleur_de_lis",
		Category: "symbols",
	},
	"〽️": &emojiMeta{
		Keywords: []string{"graph", "presentation", "stats", "business", "economics", "bad"},
		Name:     "part_alternation_mark",
		Category: "symbols",
	},
	"⚠️": &emojiMeta{
		Keywords: []string{"exclamation", "wip", "alert", "error", "problem", "issue"},
		Name:     "warning",
		Category: "symbols",
	},
	"🚸": &emojiMeta{
		Keywords: []string{"school", "warning", "danger", "sign", "driving", "yellow-diamond"},
		Name:     "children_crossing",
		Category: "symbols",
	},
	"🔰": &emojiMeta{
		Keywords: []string{"badge", "shield"},
		Name:     "beginner",
		Category: "symbols",
	},
	"♻️": &emojiMeta{
		Keywords: []string{"arrow", "environment", "garbage", "trash"},
		Name:     "recycle",
		Category: "symbols",
	},
	"🈯": &emojiMeta{
		Keywords: []string{"chinese", "point", "green-square", "kanji"},
		Name:     "u6307",
		Category: "symbols",
	},
	"💹": &emojiMeta{
		Keywords: []string{"green-square", "graph", "presentation", "stats"},
		Name:     "chart",
		Category: "symbols",
	},
	"❇️": &emojiMeta{
		Keywords: []string{"stars", "green-square", "awesome", "good", "fireworks"},
		Name:     "sparkle",
		Category: "symbols",
	},
	"✳️": &emojiMeta{
		Keywords: []string{"star", "sparkle", "green-square"},
		Name:     "eight_spoked_asterisk",
		Category: "symbols",
	},
	"❎": &emojiMeta{
		Keywords: []string{"x", "green-square", "no", "deny"},
		Name:     "negative_squared_cross_mark",
		Category: "symbols",
	},
	"✅": &emojiMeta{
		Keywords: []string{"green-square", "ok", "agree", "vote", "election", "answer", "tick"},
		Name:     "white_check_mark",
		Category: "symbols",
	},
	"💠": &emojiMeta{
		Keywords: []string{"jewel", "blue", "gem", "crystal", "fancy"},
		Name:     "diamond_shape_with_a_dot_inside",
		Category: "symbols",
	},
	"🌀": &emojiMeta{
		Keywords: []string{"weather", "swirl", "blue", "cloud", "vortex", "spiral", "whirlpool", "spin", "tornado", "hurricane", "typhoon"},
		Name:     "cyclone",
		Category: "symbols",
	},
	"➿": &emojiMeta{
		Keywords: []string{"tape", "cassette"},
		Name:     "loop",
		Category: "symbols",
	},
	"🌐": &emojiMeta{
		Keywords: []string{"earth", "international", "world", "internet", "interweb", "i18n"},
		Name:     "globe_with_meridians",
		Category: "symbols",
	},
	"Ⓜ️": &emojiMeta{
		Keywords: []string{"alphabet", "blue-circle", "letter"},
		Name:     "m",
		Category: "symbols",
	},
	"🏧": &emojiMeta{
		Keywords: []string{"money", "sales", "cash", "blue-square", "payment", "bank"},
		Name:     "atm",
		Category: "symbols",
	},
	"🈂️": &emojiMeta{
		Keywords: []string{"japanese", "blue-square", "katakana"},
		Name:     "sa",
		Category: "symbols",
	},
	"🛂": &emojiMeta{
		Keywords: []string{"custom", "blue-square"},
		Name:     "passport_control",
		Category: "symbols",
	},
	"🛃": &emojiMeta{
		Keywords: []string{"passport", "border", "blue-square"},
		Name:     "customs",
		Category: "symbols",
	},
	"🛄": &emojiMeta{
		Keywords: []string{"blue-square", "airport", "transport"},
		Name:     "baggage_claim",
		Category: "symbols",
	},
	"🛅": &emojiMeta{
		Keywords: []string{"blue-square", "travel"},
		Name:     "left_luggage",
		Category: "symbols",
	},
	"♿": &emojiMeta{
		Keywords: []string{"blue-square", "disabled", "a11y", "accessibility"},
		Name:     "wheelchair",
		Category: "symbols",
	},
	"🚭": &emojiMeta{
		Keywords: []string{"cigarette", "blue-square", "smell", "smoke"},
		Name:     "no_smoking",
		Category: "symbols",
	},
	"🚾": &emojiMeta{
		Keywords: []string{"toilet", "restroom", "blue-square"},
		Name:     "wc",
		Category: "symbols",
	},
	"🅿️": &emojiMeta{
		Keywords: []string{"cars", "blue-square", "alphabet", "letter"},
		Name:     "parking",
		Category: "symbols",
	},
	"🚰": &emojiMeta{
		Keywords: []string{"blue-square", "liquid", "restroom", "cleaning", "faucet"},
		Name:     "potable_water",
		Category: "symbols",
	},
	"🚹": &emojiMeta{
		Keywords: []string{"toilet", "restroom", "wc", "blue-square", "gender", "male"},
		Name:     "mens",
		Category: "symbols",
	},
	"🚺": &emojiMeta{
		Keywords: []string{"purple-square", "woman", "female", "toilet", "loo", "restroom", "gender"},
		Name:     "womens",
		Category: "symbols",
	},
	"🚼": &emojiMeta{
		Keywords: []string{"orange-square", "child"},
		Name:     "baby_symbol",
		Category: "symbols",
	},
	"🚻": &emojiMeta{
		Keywords: []string{"blue-square", "toilet", "refresh", "wc", "gender"},
		Name:     "restroom",
		Category: "symbols",
	},
	"🚮": &emojiMeta{
		Keywords: []string{"blue-square", "sign", "human", "info"},
		Name:     "put_litter_in_its_place",
		Category: "symbols",
	},
	"🎦": &emojiMeta{
		Keywords: []string{"blue-square", "record", "film", "movie", "curtain", "stage", "theater"},
		Name:     "cinema",
		Category: "symbols",
	},
	"📶": &emojiMeta{
		Keywords: []string{"blue-square", "reception", "phone", "internet", "connection", "wifi", "bluetooth", "bars"},
		Name:     "signal_strength",
		Category: "symbols",
	},
	"🈁": &emojiMeta{
		Keywords: []string{"blue-square", "here", "katakana", "japanese", "destination"},
		Name:     "koko",
		Category: "symbols",
	},
	"🆖": &emojiMeta{
		Keywords: []string{"blue-square", "words", "shape", "icon"},
		Name:     "ng",
		Category: "symbols",
	},
	"🆗": &emojiMeta{
		Keywords: []string{"good", "agree", "yes", "blue-square"},
		Name:     "ok",
		Category: "symbols",
	},
	"🆙": &emojiMeta{
		Keywords: []string{"blue-square", "above", "high"},
		Name:     "up",
		Category: "symbols",
	},
	"🆒": &emojiMeta{
		Keywords: []string{"words", "blue-square"},
		Name:     "cool",
		Category: "symbols",
	},
	"🆕": &emojiMeta{
		Keywords: []string{"blue-square", "words", "start"},
		Name:     "new",
		Category: "symbols",
	},
	"🆓": &emojiMeta{
		Keywords: []string{"blue-square", "words"},
		Name:     "free",
		Category: "symbols",
	},
	"0️⃣": &emojiMeta{
		Keywords: []string{"0", "numbers", "blue-square", "null"},
		Name:     "zero",
		Category: "symbols",
	},
	"1️⃣": &emojiMeta{
		Keywords: []string{"blue-square", "numbers", "1"},
		Name:     "one",
		Category: "symbols",
	},
	"2️⃣": &emojiMeta{
		Keywords: []string{"numbers", "2", "prime", "blue-square"},
		Name:     "two",
		Category: "symbols",
	},
	"3️⃣": &emojiMeta{
		Keywords: []string{"3", "numbers", "prime", "blue-square"},
		Name:     "three",
		Category: "symbols",
	},
	"4️⃣": &emojiMeta{
		Keywords: []string{"4", "numbers", "blue-square"},
		Name:     "four",
		Category: "symbols",
	},
	"5️⃣": &emojiMeta{
		Keywords: []string{"5", "numbers", "blue-square", "prime"},
		Name:     "five",
		Category: "symbols",
	},
	"6️⃣": &emojiMeta{
		Keywords: []string{"6", "numbers", "blue-square"},
		Name:     "six",
		Category: "symbols",
	},
	"7️⃣": &emojiMeta{
		Keywords: []string{"7", "numbers", "blue-square", "prime"},
		Name:     "seven",
		Category: "symbols",
	},
	"8️⃣": &emojiMeta{
		Keywords: []string{"8", "blue-square", "numbers"},
		Name:     "eight",
		Category: "symbols",
	},
	"9️⃣": &emojiMeta{
		Keywords: []string{"blue-square", "numbers", "9"},
		Name:     "nine",
		Category: "symbols",
	},
	"🔟": &emojiMeta{
		Keywords: []string{"numbers", "10", "blue-square"},
		Name:     "keycap_ten",
		Category: "symbols",
	},
	"*⃣": &emojiMeta{
		Keywords: []string{"star", "keycap"},
		Name:     "asterisk",
		Category: "symbols",
	},
	"🔢": &emojiMeta{
		Keywords: []string{"numbers", "blue-square"},
		Name:     "1234",
		Category: "symbols",
	},
	"⏏️": &emojiMeta{
		Keywords: []string{"blue-square"},
		Name:     "eject_button",
		Category: "symbols",
	},
	"▶️": &emojiMeta{
		Keywords: []string{"blue-square", "right", "direction", "play"},
		Name:     "arrow_forward",
		Category: "symbols",
	},
	"⏸": &emojiMeta{
		Keywords: []string{"pause", "blue-square"},
		Name:     "pause_button",
		Category: "symbols",
	},
	"⏭": &emojiMeta{
		Keywords: []string{"forward", "next", "blue-square"},
		Name:     "next_track_button",
		Category: "symbols",
	},
	"⏹": &emojiMeta{
		Keywords: []string{"blue-square"},
		Name:     "stop_button",
		Category: "symbols",
	},
	"⏺": &emojiMeta{
		Keywords: []string{"blue-square"},
		Name:     "record_button",
		Category: "symbols",
	},
	"⏯": &emojiMeta{
		Keywords: []string{"blue-square", "play", "pause"},
		Name:     "play_or_pause_button",
		Category: "symbols",
	},
	"⏮": &emojiMeta{
		Keywords: []string{"backward"},
		Name:     "previous_track_button",
		Category: "symbols",
	},
	"⏩": &emojiMeta{
		Keywords: []string{"blue-square", "play", "speed", "continue"},
		Name:     "fast_forward",
		Category: "symbols",
	},
	"⏪": &emojiMeta{
		Keywords: []string{"play", "blue-square"},
		Name:     "rewind",
		Category: "symbols",
	},
	"🔀": &emojiMeta{
		Keywords: []string{"blue-square", "shuffle", "music", "random"},
		Name:     "twisted_rightwards_arrows",
		Category: "symbols",
	},
	"🔁": &emojiMeta{
		Keywords: []string{"loop", "record"},
		Name:     "repeat",
		Category: "symbols",
	},
	"🔂": &emojiMeta{
		Keywords: []string{"blue-square", "loop"},
		Name:     "repeat_one",
		Category: "symbols",
	},
	"◀️": &emojiMeta{
		Keywords: []string{"blue-square", "left", "direction"},
		Name:     "arrow_backward",
		Category: "symbols",
	},
	"🔼": &emojiMeta{
		Keywords: []string{"blue-square", "triangle", "direction", "point", "forward", "top"},
		Name:     "arrow_up_small",
		Category: "symbols",
	},
	"🔽": &emojiMeta{
		Keywords: []string{"blue-square", "direction", "bottom"},
		Name:     "arrow_down_small",
		Category: "symbols",
	},
	"⏫": &emojiMeta{
		Keywords: []string{"blue-square", "direction", "top"},
		Name:     "arrow_double_up",
		Category: "symbols",
	},
	"⏬": &emojiMeta{
		Keywords: []string{"blue-square", "direction", "bottom"},
		Name:     "arrow_double_down",
		Category: "symbols",
	},
	"➡️": &emojiMeta{
		Keywords: []string{"blue-square", "next"},
		Name:     "arrow_right",
		Category: "symbols",
	},
	"⬅️": &emojiMeta{
		Keywords: []string{"blue-square", "previous", "back"},
		Name:     "arrow_left",
		Category: "symbols",
	},
	"⬆️": &emojiMeta{
		Keywords: []string{"blue-square", "continue", "top", "direction"},
		Name:     "arrow_up",
		Category: "symbols",
	},
	"⬇️": &emojiMeta{
		Keywords: []string{"blue-square", "direction", "bottom"},
		Name:     "arrow_down",
		Category: "symbols",
	},
	"↗️": &emojiMeta{
		Keywords: []string{"blue-square", "point", "direction", "diagonal", "northeast"},
		Name:     "arrow_upper_right",
		Category: "symbols",
	},
	"↘️": &emojiMeta{
		Keywords: []string{"blue-square", "direction", "diagonal", "southeast"},
		Name:     "arrow_lower_right",
		Category: "symbols",
	},
	"↙️": &emojiMeta{
		Keywords: []string{"blue-square", "direction", "diagonal", "southwest"},
		Name:     "arrow_lower_left",
		Category: "symbols",
	},
	"↖️": &emojiMeta{
		Keywords: []string{"blue-square", "point", "direction", "diagonal", "northwest"},
		Name:     "arrow_upper_left",
		Category: "symbols",
	},
	"↕️": &emojiMeta{
		Keywords: []string{"blue-square", "direction", "way", "vertical"},
		Name:     "arrow_up_down",
		Category: "symbols",
	},
	"↔️": &emojiMeta{
		Keywords: []string{"shape", "direction", "horizontal", "sideways"},
		Name:     "left_right_arrow",
		Category: "symbols",
	},
	"🔄": &emojiMeta{
		Keywords: []string{"blue-square", "sync", "cycle"},
		Name:     "arrows_counterclockwise",
		Category: "symbols",
	},
	"↪️": &emojiMeta{
		Keywords: []string{"blue-square", "return", "rotate", "direction"},
		Name:     "arrow_right_hook",
		Category: "symbols",
	},
	"↩️": &emojiMeta{
		Keywords: []string{"back", "return", "blue-square", "undo", "enter"},
		Name:     "leftwards_arrow_with_hook",
		Category: "symbols",
	},
	"⤴️": &emojiMeta{
		Keywords: []string{"blue-square", "direction", "top"},
		Name:     "arrow_heading_up",
		Category: "symbols",
	},
	"⤵️": &emojiMeta{
		Keywords: []string{"blue-square", "direction", "bottom"},
		Name:     "arrow_heading_down",
		Category: "symbols",
	},
	"#️⃣": &emojiMeta{
		Keywords: []string{"symbol", "blue-square", "twitter"},
		Name:     "hash",
		Category: "symbols",
	},
	"ℹ️": &emojiMeta{
		Keywords: []string{"blue-square", "alphabet", "letter"},
		Name:     "information_source",
		Category: "symbols",
	},
	"🔤": &emojiMeta{
		Keywords: []string{"blue-square", "alphabet"},
		Name:     "abc",
		Category: "symbols",
	},
	"🔡": &emojiMeta{
		Keywords: []string{"blue-square", "alphabet"},
		Name:     "abcd",
		Category: "symbols",
	},
	"🔠": &emojiMeta{
		Keywords: []string{"alphabet", "words", "blue-square"},
		Name:     "capital_abcd",
		Category: "symbols",
	},
	"🔣": &emojiMeta{
		Keywords: []string{"blue-square", "music", "note", "ampersand", "percent", "glyphs", "characters"},
		Name:     "symbols",
		Category: "symbols",
	},
	"🎵": &emojiMeta{
		Keywords: []string{"score", "tone", "sound"},
		Name:     "musical_note",
		Category: "symbols",
	},
	"🎶": &emojiMeta{
		Keywords: []string{"music", "score"},
		Name:     "notes",
		Category: "symbols",
	},
	"〰️": &emojiMeta{
		Keywords: []string{"draw", "line", "moustache", "mustache", "squiggle", "scribble"},
		Name:     "wavy_dash",
		Category: "symbols",
	},
	"➰": &emojiMeta{
		Keywords: []string{"scribble", "draw", "shape", "squiggle"},
		Name:     "curly_loop",
		Category: "symbols",
	},
	"✔️": &emojiMeta{
		Keywords: []string{"ok", "nike", "answer", "yes", "tick"},
		Name:     "heavy_check_mark",
		Category: "symbols",
	},
	"🔃": &emojiMeta{
		Keywords: []string{"sync", "cycle", "round", "repeat"},
		Name:     "arrows_clockwise",
		Category: "symbols",
	},
	"➕": &emojiMeta{
		Keywords: []string{"math", "calculation", "addition", "more", "increase"},
		Name:     "heavy_plus_sign",
		Category: "symbols",
	},
	"➖": &emojiMeta{
		Keywords: []string{"math", "calculation", "subtract", "less"},
		Name:     "heavy_minus_sign",
		Category: "symbols",
	},
	"➗": &emojiMeta{
		Keywords: []string{"divide", "math", "calculation"},
		Name:     "heavy_division_sign",
		Category: "symbols",
	},
	"✖️": &emojiMeta{
		Keywords: []string{"math", "calculation"},
		Name:     "heavy_multiplication_x",
		Category: "symbols",
	},
	"♾": &emojiMeta{
		Keywords: []string{"forever"},
		Name:     "infinity",
		Category: "symbols",
	},
	"💲": &emojiMeta{
		Keywords: []string{"money", "sales", "payment", "currency", "buck"},
		Name:     "heavy_dollar_sign",
		Category: "symbols",
	},
	"💱": &emojiMeta{
		Keywords: []string{"money", "sales", "dollar", "travel"},
		Name:     "currency_exchange",
		Category: "symbols",
	},
	"©️": &emojiMeta{
		Keywords: []string{"ip", "license", "circle", "law", "legal"},
		Name:     "copyright",
		Category: "symbols",
	},
	"®️": &emojiMeta{
		Keywords: []string{"alphabet", "circle"},
		Name:     "registered",
		Category: "symbols",
	},
	"™️": &emojiMeta{
		Keywords: []string{"trademark", "brand", "law", "legal"},
		Name:     "tm",
		Category: "symbols",
	},
	"🔚": &emojiMeta{
		Keywords: []string{"words", "arrow"},
		Name:     "end",
		Category: "symbols",
	},
	"🔙": &emojiMeta{
		Keywords: []string{"arrow", "words", "return"},
		Name:     "back",
		Category: "symbols",
	},
	"🔛": &emojiMeta{
		Keywords: []string{"arrow", "words"},
		Name:     "on",
		Category: "symbols",
	},
	"🔝": &emojiMeta{
		Keywords: []string{"words", "blue-square"},
		Name:     "top",
		Category: "symbols",
	},
	"🔜": &emojiMeta{
		Keywords: []string{"arrow", "words"},
		Name:     "soon",
		Category: "symbols",
	},
	"☑️": &emojiMeta{
		Keywords: []string{"ok", "agree", "confirm", "black-square", "vote", "election", "yes", "tick"},
		Name:     "ballot_box_with_check",
		Category: "symbols",
	},
	"🔘": &emojiMeta{
		Keywords: []string{"input", "old", "music", "circle"},
		Name:     "radio_button",
		Category: "symbols",
	},
	"⚪": &emojiMeta{
		Keywords: []string{"shape", "round"},
		Name:     "white_circle",
		Category: "symbols",
	},
	"⚫": &emojiMeta{
		Keywords: []string{"shape", "button", "round"},
		Name:     "black_circle",
		Category: "symbols",
	},
	"🔴": &emojiMeta{
		Keywords: []string{"shape", "error", "danger"},
		Name:     "red_circle",
		Category: "symbols",
	},
	"🔵": &emojiMeta{
		Keywords: []string{"shape", "icon", "button"},
		Name:     "large_blue_circle",
		Category: "symbols",
	},
	"🔸": &emojiMeta{
		Keywords: []string{"shape", "jewel", "gem"},
		Name:     "small_orange_diamond",
		Category: "symbols",
	},
	"🔹": &emojiMeta{
		Keywords: []string{"shape", "jewel", "gem"},
		Name:     "small_blue_diamond",
		Category: "symbols",
	},
	"🔶": &emojiMeta{
		Keywords: []string{"shape", "jewel", "gem"},
		Name:     "large_orange_diamond",
		Category: "symbols",
	},
	"🔷": &emojiMeta{
		Keywords: []string{"shape", "jewel", "gem"},
		Name:     "large_blue_diamond",
		Category: "symbols",
	},
	"🔺": &emojiMeta{
		Keywords: []string{"shape", "direction", "up", "top"},
		Name:     "small_red_triangle",
		Category: "symbols",
	},
	"▪️": &emojiMeta{
		Keywords: []string{"shape", "icon"},
		Name:     "black_small_square",
		Category: "symbols",
	},
	"▫️": &emojiMeta{
		Keywords: []string{"shape", "icon"},
		Name:     "white_small_square",
		Category: "symbols",
	},
	"⬛": &emojiMeta{
		Keywords: []string{"shape", "icon", "button"},
		Name:     "black_large_square",
		Category: "symbols",
	},
	"⬜": &emojiMeta{
		Keywords: []string{"shape", "icon", "stone", "button"},
		Name:     "white_large_square",
		Category: "symbols",
	},
	"🔻": &emojiMeta{
		Keywords: []string{"shape", "direction", "bottom"},
		Name:     "small_red_triangle_down",
		Category: "symbols",
	},
	"◼️": &emojiMeta{
		Keywords: []string{"shape", "button", "icon"},
		Name:     "black_medium_square",
		Category: "symbols",
	},
	"◻️": &emojiMeta{
		Keywords: []string{"shape", "stone", "icon"},
		Name:     "white_medium_square",
		Category: "symbols",
	},
	"◾": &emojiMeta{
		Keywords: []string{"icon", "shape", "button"},
		Name:     "black_medium_small_square",
		Category: "symbols",
	},
	"◽": &emojiMeta{
		Keywords: []string{"shape", "stone", "icon", "button"},
		Name:     "white_medium_small_square",
		Category: "symbols",
	},
	"🔲": &emojiMeta{
		Keywords: []string{"shape", "input", "frame"},
		Name:     "black_square_button",
		Category: "symbols",
	},
	"🔳": &emojiMeta{
		Keywords: []string{"shape", "input"},
		Name:     "white_square_button",
		Category: "symbols",
	},
	"🔈": &emojiMeta{
		Keywords: []string{"sound", "volume", "silence", "broadcast"},
		Name:     "speaker",
		Category: "symbols",
	},
	"🔉": &emojiMeta{
		Keywords: []string{"volume", "speaker", "broadcast"},
		Name:     "sound",
		Category: "symbols",
	},
	"🔊": &emojiMeta{
		Keywords: []string{"volume", "noise", "noisy", "speaker", "broadcast"},
		Name:     "loud_sound",
		Category: "symbols",
	},
	"🔇": &emojiMeta{
		Keywords: []string{"sound", "volume", "silence", "quiet"},
		Name:     "mute",
		Category: "symbols",
	},
	"📣": &emojiMeta{
		Keywords: []string{"sound", "speaker", "volume"},
		Name:     "mega",
		Category: "symbols",
	},
	"📢": &emojiMeta{
		Keywords: []string{"volume", "sound"},
		Name:     "loudspeaker",
		Category: "symbols",
	},
	"🔔": &emojiMeta{
		Keywords: []string{"sound", "notification", "christmas", "xmas", "chime"},
		Name:     "bell",
		Category: "symbols",
	},
	"🔕": &emojiMeta{
		Keywords: []string{"sound", "volume", "mute", "quiet", "silent"},
		Name:     "no_bell",
		Category: "symbols",
	},
	"🃏": &emojiMeta{
		Keywords: []string{"poker", "cards", "game", "play", "magic"},
		Name:     "black_joker",
		Category: "symbols",
	},
	"🀄": &emojiMeta{
		Keywords: []string{"game", "play", "chinese", "kanji"},
		Name:     "mahjong",
		Category: "symbols",
	},
	"♠️": &emojiMeta{
		Keywords: []string{"poker", "cards", "suits", "magic"},
		Name:     "spades",
		Category: "symbols",
	},
	"♣️": &emojiMeta{
		Keywords: []string{"poker", "cards", "magic", "suits"},
		Name:     "clubs",
		Category: "symbols",
	},
	"♥️": &emojiMeta{
		Keywords: []string{"poker", "cards", "magic", "suits"},
		Name:     "hearts",
		Category: "symbols",
	},
	"♦️": &emojiMeta{
		Keywords: []string{"poker", "cards", "magic", "suits"},
		Name:     "diamonds",
		Category: "symbols",
	},
	"🎴": &emojiMeta{
		Keywords: []string{"game", "sunset", "red"},
		Name:     "flower_playing_cards",
		Category: "symbols",
	},
	"💭": &emojiMeta{
		Keywords: []string{"bubble", "cloud", "speech", "thinking", "dream"},
		Name:     "thought_balloon",
		Category: "symbols",
	},
	"🗯": &emojiMeta{
		Keywords: []string{"caption", "speech", "thinking", "mad"},
		Name:     "right_anger_bubble",
		Category: "symbols",
	},
	"💬": &emojiMeta{
		Keywords: []string{"bubble", "words", "message", "talk", "chatting"},
		Name:     "speech_balloon",
		Category: "symbols",
	},
	"🗨": &emojiMeta{
		Keywords: []string{"words", "message", "talk", "chatting"},
		Name:     "left_speech_bubble",
		Category: "symbols",
	},
	"🕐": &emojiMeta{
		Keywords: []string{"time", "late", "early", "schedule"},
		Name:     "clock1",
		Category: "symbols",
	},
	"🕑": &emojiMeta{
		Keywords: []string{"time", "late", "early", "schedule"},
		Name:     "clock2",
		Category: "symbols",
	},
	"🕒": &emojiMeta{
		Keywords: []string{"time", "late", "early", "schedule"},
		Name:     "clock3",
		Category: "symbols",
	},
	"🕓": &emojiMeta{
		Keywords: []string{"time", "late", "early", "schedule"},
		Name:     "clock4",
		Category: "symbols",
	},
	"🕔": &emojiMeta{
		Keywords: []string{"time", "late", "early", "schedule"},
		Name:     "clock5",
		Category: "symbols",
	},
	"🕕": &emojiMeta{
		Keywords: []string{"time", "late", "early", "schedule", "dawn", "dusk"},
		Name:     "clock6",
		Category: "symbols",
	},
	"🕖": &emojiMeta{
		Keywords: []string{"time", "late", "early", "schedule"},
		Name:     "clock7",
		Category: "symbols",
	},
	"🕗": &emojiMeta{
		Keywords: []string{"time", "late", "early", "schedule"},
		Name:     "clock8",
		Category: "symbols",
	},
	"🕘": &emojiMeta{
		Keywords: []string{"time", "late", "early", "schedule"},
		Name:     "clock9",
		Category: "symbols",
	},
	"🕙": &emojiMeta{
		Keywords: []string{"time", "late", "early", "schedule"},
		Name:     "clock10",
		Category: "symbols",
	},
	"🕚": &emojiMeta{
		Keywords: []string{"time", "late", "early", "schedule"},
		Name:     "clock11",
		Category: "symbols",
	},
	"🕛": &emojiMeta{
		Keywords: []string{"time", "noon", "midnight", "midday", "late", "early", "schedule"},
		Name:     "clock12",
		Category: "symbols",
	},
	"🕜": &emojiMeta{
		Keywords: []string{"time", "late", "early", "schedule"},
		Name:     "clock130",
		Category: "symbols",
	},
	"🕝": &emojiMeta{
		Keywords: []string{"time", "late", "early", "schedule"},
		Name:     "clock230",
		Category: "symbols",
	},
	"🕞": &emojiMeta{
		Keywords: []string{"time", "late", "early", "schedule"},
		Name:     "clock330",
		Category: "symbols",
	},
	"🕟": &emojiMeta{
		Keywords: []string{"time", "late", "early", "schedule"},
		Name:     "clock430",
		Category: "symbols",
	},
	"🕠": &emojiMeta{
		Keywords: []string{"time", "late", "early", "schedule"},
		Name:     "clock530",
		Category: "symbols",
	},
	"🕡": &emojiMeta{
		Keywords: []string{"time", "late", "early", "schedule"},
		Name:     "clock630",
		Category: "symbols",
	},
	"🕢": &emojiMeta{
		Keywords: []string{"time", "late", "early", "schedule"},
		Name:     "clock730",
		Category: "symbols",
	},
	"🕣": &emojiMeta{
		Keywords: []string{"time", "late", "early", "schedule"},
		Name:     "clock830",
		Category: "symbols",
	},
	"🕤": &emojiMeta{
		Keywords: []string{"time", "late", "early", "schedule"},
		Name:     "clock930",
		Category: "symbols",
	},
	"🕥": &emojiMeta{
		Keywords: []string{"time", "late", "early", "schedule"},
		Name:     "clock1030",
		Category: "symbols",
	},
	"🕦": &emojiMeta{
		Keywords: []string{"time", "late", "early", "schedule"},
		Name:     "clock1130",
		Category: "symbols",
	},
	"🕧": &emojiMeta{
		Keywords: []string{"time", "late", "early", "schedule"},
		Name:     "clock1230",
		Category: "symbols",
	},
	"🇦🇫": &emojiMeta{
		Keywords: []string{"af", "flag", "nation", "country", "banner"},
		Name:     "afghanistan",
		Category: "flag",
	},
	"🇦🇽": &emojiMeta{
		Keywords: []string{"Åland", "islands", "flag", "nation", "country", "banner"},
		Name:     "aland_islands",
		Category: "flag",
	},
	"🇦🇱": &emojiMeta{
		Keywords: []string{"al", "flag", "nation", "country", "banner"},
		Name:     "albania",
		Category: "flag",
	},
	"🇩🇿": &emojiMeta{
		Keywords: []string{"dz", "flag", "nation", "country", "banner"},
		Name:     "algeria",
		Category: "flag",
	},
	"🇦🇸": &emojiMeta{
		Keywords: []string{"american", "ws", "flag", "nation", "country", "banner"},
		Name:     "american_samoa",
		Category: "flag",
	},
	"🇦🇩": &emojiMeta{
		Keywords: []string{"ad", "flag", "nation", "country", "banner"},
		Name:     "andorra",
		Category: "flag",
	},
	"🇦🇴": &emojiMeta{
		Keywords: []string{"ao", "flag", "nation", "country", "banner"},
		Name:     "angola",
		Category: "flag",
	},
	"🇦🇮": &emojiMeta{
		Keywords: []string{"ai", "flag", "nation", "country", "banner"},
		Name:     "anguilla",
		Category: "flag",
	},
	"🇦🇶": &emojiMeta{
		Keywords: []string{"aq", "flag", "nation", "country", "banner"},
		Name:     "antarctica",
		Category: "flag",
	},
	"🇦🇬": &emojiMeta{
		Keywords: []string{"antigua", "barbuda", "flag", "nation", "country", "banner"},
		Name:     "antigua_barbuda",
		Category: "flag",
	},
	"🇦🇷": &emojiMeta{
		Keywords: []string{"ar", "flag", "nation", "country", "banner"},
		Name:     "argentina",
		Category: "flag",
	},
	"🇦🇲": &emojiMeta{
		Keywords: []string{"am", "flag", "nation", "country", "banner"},
		Name:     "armenia",
		Category: "flag",
	},
	"🇦🇼": &emojiMeta{
		Keywords: []string{"aw", "flag", "nation", "country", "banner"},
		Name:     "aruba",
		Category: "flag",
	},
	"🇦🇺": &emojiMeta{
		Keywords: []string{"au", "flag", "nation", "country", "banner"},
		Name:     "australia",
		Category: "flag",
	},
	"🇦🇹": &emojiMeta{
		Keywords: []string{"at", "flag", "nation", "country", "banner"},
		Name:     "austria",
		Category: "flag",
	},
	"🇦🇿": &emojiMeta{
		Keywords: []string{"az", "flag", "nation", "country", "banner"},
		Name:     "azerbaijan",
		Category: "flag",
	},
	"🇧🇸": &emojiMeta{
		Keywords: []string{"bs", "flag", "nation", "country", "banner"},
		Name:     "bahamas",
		Category: "flag",
	},
	"🇧🇭": &emojiMeta{
		Keywords: []string{"bh", "flag", "nation", "country", "banner"},
		Name:     "bahrain",
		Category: "flag",
	},
	"🇧🇩": &emojiMeta{
		Keywords: []string{"bd", "flag", "nation", "country", "banner"},
		Name:     "bangladesh",
		Category: "flag",
	},
	"🇧🇧": &emojiMeta{
		Keywords: []string{"bb", "flag", "nation", "country", "banner"},
		Name:     "barbados",
		Category: "flag",
	},
	"🇧🇾": &emojiMeta{
		Keywords: []string{"by", "flag", "nation", "country", "banner"},
		Name:     "belarus",
		Category: "flag",
	},
	"🇧🇪": &emojiMeta{
		Keywords: []string{"be", "flag", "nation", "country", "banner"},
		Name:     "belgium",
		Category: "flag",
	},
	"🇧🇿": &emojiMeta{
		Keywords: []string{"bz", "flag", "nation", "country", "banner"},
		Name:     "belize",
		Category: "flag",
	},
	"🇧🇯": &emojiMeta{
		Keywords: []string{"bj", "flag", "nation", "country", "banner"},
		Name:     "benin",
		Category: "flag",
	},
	"🇧🇲": &emojiMeta{
		Keywords: []string{"bm", "flag", "nation", "country", "banner"},
		Name:     "bermuda",
		Category: "flag",
	},
	"🇧🇹": &emojiMeta{
		Keywords: []string{"bt", "flag", "nation", "country", "banner"},
		Name:     "bhutan",
		Category: "flag",
	},
	"🇧🇴": &emojiMeta{
		Keywords: []string{"bo", "flag", "nation", "country", "banner"},
		Name:     "bolivia",
		Category: "flag",
	},
	"🇧🇶": &emojiMeta{
		Keywords: []string{"bonaire", "flag", "nation", "country", "banner"},
		Name:     "caribbean_netherlands",
		Category: "flag",
	},
	"🇧🇦": &emojiMeta{
		Keywords: []string{"bosnia", "herzegovina", "flag", "nation", "country", "banner"},
		Name:     "bosnia_herzegovina",
		Category: "flag",
	},
	"🇧🇼": &emojiMeta{
		Keywords: []string{"bw", "flag", "nation", "country", "banner"},
		Name:     "botswana",
		Category: "flag",
	},
	"🇧🇷": &emojiMeta{
		Keywords: []string{"br", "flag", "nation", "country", "banner"},
		Name:     "brazil",
		Category: "flag",
	},
	"🇮🇴": &emojiMeta{
		Keywords: []string{"british", "indian", "ocean", "territory", "flag", "nation", "country", "banner"},
		Name:     "british_indian_ocean_territory",
		Category: "flag",
	},
	"🇻🇬": &emojiMeta{
		Keywords: []string{"british", "virgin", "islands", "bvi", "flag", "nation", "country", "banner"},
		Name:     "british_virgin_islands",
		Category: "flag",
	},
	"🇧🇳": &emojiMeta{
		Keywords: []string{"bn", "darussalam", "flag", "nation", "country", "banner"},
		Name:     "brunei",
		Category: "flag",
	},
	"🇧🇬": &emojiMeta{
		Keywords: []string{"bg", "flag", "nation", "country", "banner"},
		Name:     "bulgaria",
		Category: "flag",
	},
	"🇧🇫": &emojiMeta{
		Keywords: []string{"burkina", "faso", "flag", "nation", "country", "banner"},
		Name:     "burkina_faso",
		Category: "flag",
	},
	"🇧🇮": &emojiMeta{
		Keywords: []string{"bi", "flag", "nation", "country", "banner"},
		Name:     "burundi",
		Category: "flag",
	},
	"🇨🇻": &emojiMeta{
		Keywords: []string{"cabo", "verde", "flag", "nation", "country", "banner"},
		Name:     "cape_verde",
		Category: "flag",
	},
	"🇰🇭": &emojiMeta{
		Keywords: []string{"kh", "flag", "nation", "country", "banner"},
		Name:     "cambodia",
		Category: "flag",
	},
	"🇨🇲": &emojiMeta{
		Keywords: []string{"cm", "flag", "nation", "country", "banner"},
		Name:     "cameroon",
		Category: "flag",
	},
	"🇨🇦": &emojiMeta{
		Keywords: []string{"ca", "flag", "nation", "country", "banner"},
		Name:     "canada",
		Category: "flag",
	},
	"🇮🇨": &emojiMeta{
		Keywords: []string{"canary", "islands", "flag", "nation", "country", "banner"},
		Name:     "canary_islands",
		Category: "flag",
	},
	"🇰🇾": &emojiMeta{
		Keywords: []string{"cayman", "islands", "flag", "nation", "country", "banner"},
		Name:     "cayman_islands",
		Category: "flag",
	},
	"🇨🇫": &emojiMeta{
		Keywords: []string{"central", "african", "republic", "flag", "nation", "country", "banner"},
		Name:     "central_african_republic",
		Category: "flag",
	},
	"🇹🇩": &emojiMeta{
		Keywords: []string{"td", "flag", "nation", "country", "banner"},
		Name:     "chad",
		Category: "flag",
	},
	"🇨🇱": &emojiMeta{
		Keywords: []string{"flag", "nation", "country", "banner"},
		Name:     "chile",
		Category: "flag",
	},
	"🇨🇳": &emojiMeta{
		Keywords: []string{"china", "chinese", "prc", "flag", "country", "nation", "banner"},
		Name:     "cn",
		Category: "flag",
	},
	"🇨🇽": &emojiMeta{
		Keywords: []string{"christmas", "island", "flag", "nation", "country", "banner"},
		Name:     "christmas_island",
		Category: "flag",
	},
	"🇨🇨": &emojiMeta{
		Keywords: []string{"cocos", "keeling", "islands", "flag", "nation", "country", "banner"},
		Name:     "cocos_islands",
		Category: "flag",
	},
	"🇨🇴": &emojiMeta{
		Keywords: []string{"co", "flag", "nation", "country", "banner"},
		Name:     "colombia",
		Category: "flag",
	},
	"🇰🇲": &emojiMeta{
		Keywords: []string{"km", "flag", "nation", "country", "banner"},
		Name:     "comoros",
		Category: "flag",
	},
	"🇨🇬": &emojiMeta{
		Keywords: []string{"congo", "flag", "nation", "country", "banner"},
		Name:     "congo_brazzaville",
		Category: "flag",
	},
	"🇨🇩": &emojiMeta{
		Keywords: []string{"congo", "democratic", "republic", "flag", "nation", "country", "banner"},
		Name:     "congo_kinshasa",
		Category: "flag",
	},
	"🇨🇰": &emojiMeta{
		Keywords: []string{"cook", "islands", "flag", "nation", "country", "banner"},
		Name:     "cook_islands",
		Category: "flag",
	},
	"🇨🇷": &emojiMeta{
		Keywords: []string{"costa", "rica", "flag", "nation", "country", "banner"},
		Name:     "costa_rica",
		Category: "flag",
	},
	"🇭🇷": &emojiMeta{
		Keywords: []string{"hr", "flag", "nation", "country", "banner"},
		Name:     "croatia",
		Category: "flag",
	},
	"🇨🇺": &emojiMeta{
		Keywords: []string{"cu", "flag", "nation", "country", "banner"},
		Name:     "cuba",
		Category: "flag",
	},
	"🇨🇼": &emojiMeta{
		Keywords: []string{"curaçao", "flag", "nation", "country", "banner"},
		Name:     "curacao",
		Category: "flag",
	},
	"🇨🇾": &emojiMeta{
		Keywords: []string{"cy", "flag", "nation", "country", "banner"},
		Name:     "cyprus",
		Category: "flag",
	},
	"🇨🇿": &emojiMeta{
		Keywords: []string{"cz", "flag", "nation", "country", "banner"},
		Name:     "czech_republic",
		Category: "flag",
	},
	"🇩🇰": &emojiMeta{
		Keywords: []string{"dk", "flag", "nation", "country", "banner"},
		Name:     "denmark",
		Category: "flag",
	},
	"🇩🇯": &emojiMeta{
		Keywords: []string{"dj", "flag", "nation", "country", "banner"},
		Name:     "djibouti",
		Category: "flag",
	},
	"🇩🇲": &emojiMeta{
		Keywords: []string{"dm", "flag", "nation", "country", "banner"},
		Name:     "dominica",
		Category: "flag",
	},
	"🇩🇴": &emojiMeta{
		Keywords: []string{"dominican", "republic", "flag", "nation", "country", "banner"},
		Name:     "dominican_republic",
		Category: "flag",
	},
	"🇪🇨": &emojiMeta{
		Keywords: []string{"ec", "flag", "nation", "country", "banner"},
		Name:     "ecuador",
		Category: "flag",
	},
	"🇪🇬": &emojiMeta{
		Keywords: []string{"eg", "flag", "nation", "country", "banner"},
		Name:     "egypt",
		Category: "flag",
	},
	"🇸🇻": &emojiMeta{
		Keywords: []string{"el", "salvador", "flag", "nation", "country", "banner"},
		Name:     "el_salvador",
		Category: "flag",
	},
	"🇬🇶": &emojiMeta{
		Keywords: []string{"equatorial", "gn", "flag", "nation", "country", "banner"},
		Name:     "equatorial_guinea",
		Category: "flag",
	},
	"🇪🇷": &emojiMeta{
		Keywords: []string{"er", "flag", "nation", "country", "banner"},
		Name:     "eritrea",
		Category: "flag",
	},
	"🇪🇪": &emojiMeta{
		Keywords: []string{"ee", "flag", "nation", "country", "banner"},
		Name:     "estonia",
		Category: "flag",
	},
	"🇪🇹": &emojiMeta{
		Keywords: []string{"et", "flag", "nation", "country", "banner"},
		Name:     "ethiopia",
		Category: "flag",
	},
	"🇪🇺": &emojiMeta{
		Keywords: []string{"european", "union", "flag", "banner"},
		Name:     "eu",
		Category: "flag",
	},
	"🇫🇰": &emojiMeta{
		Keywords: []string{"falkland", "islands", "malvinas", "flag", "nation", "country", "banner"},
		Name:     "falkland_islands",
		Category: "flag",
	},
	"🇫🇴": &emojiMeta{
		Keywords: []string{"faroe", "islands", "flag", "nation", "country", "banner"},
		Name:     "faroe_islands",
		Category: "flag",
	},
	"🇫🇯": &emojiMeta{
		Keywords: []string{"fj", "flag", "nation", "country", "banner"},
		Name:     "fiji",
		Category: "flag",
	},
	"🇫🇮": &emojiMeta{
		Keywords: []string{"fi", "flag", "nation", "country", "banner"},
		Name:     "finland",
		Category: "flag",
	},
	"🇫🇷": &emojiMeta{
		Keywords: []string{"banner", "flag", "nation", "france", "french", "country"},
		Name:     "fr",
		Category: "flag",
	},
	"🇬🇫": &emojiMeta{
		Keywords: []string{"french", "guiana", "flag", "nation", "country", "banner"},
		Name:     "french_guiana",
		Category: "flag",
	},
	"🇵🇫": &emojiMeta{
		Keywords: []string{"french", "polynesia", "flag", "nation", "country", "banner"},
		Name:     "french_polynesia",
		Category: "flag",
	},
	"🇹🇫": &emojiMeta{
		Keywords: []string{"french", "southern", "territories", "flag", "nation", "country", "banner"},
		Name:     "french_southern_territories",
		Category: "flag",
	},
	"🇬🇦": &emojiMeta{
		Keywords: []string{"ga", "flag", "nation", "country", "banner"},
		Name:     "gabon",
		Category: "flag",
	},
	"🇬🇲": &emojiMeta{
		Keywords: []string{"gm", "flag", "nation", "country", "banner"},
		Name:     "gambia",
		Category: "flag",
	},
	"🇬🇪": &emojiMeta{
		Keywords: []string{"ge", "flag", "nation", "country", "banner"},
		Name:     "georgia",
		Category: "flag",
	},
	"🇩🇪": &emojiMeta{
		Keywords: []string{"german", "nation", "flag", "country", "banner"},
		Name:     "de",
		Category: "flag",
	},
	"🇬🇭": &emojiMeta{
		Keywords: []string{"gh", "flag", "nation", "country", "banner"},
		Name:     "ghana",
		Category: "flag",
	},
	"🇬🇮": &emojiMeta{
		Keywords: []string{"gi", "flag", "nation", "country", "banner"},
		Name:     "gibraltar",
		Category: "flag",
	},
	"🇬🇷": &emojiMeta{
		Keywords: []string{"gr", "flag", "nation", "country", "banner"},
		Name:     "greece",
		Category: "flag",
	},
	"🇬🇱": &emojiMeta{
		Keywords: []string{"gl", "flag", "nation", "country", "banner"},
		Name:     "greenland",
		Category: "flag",
	},
	"🇬🇩": &emojiMeta{
		Keywords: []string{"gd", "flag", "nation", "country", "banner"},
		Name:     "grenada",
		Category: "flag",
	},
	"🇬🇵": &emojiMeta{
		Keywords: []string{"gp", "flag", "nation", "country", "banner"},
		Name:     "guadeloupe",
		Category: "flag",
	},
	"🇬🇺": &emojiMeta{
		Keywords: []string{"gu", "flag", "nation", "country", "banner"},
		Name:     "guam",
		Category: "flag",
	},
	"🇬🇹": &emojiMeta{
		Keywords: []string{"gt", "flag", "nation", "country", "banner"},
		Name:     "guatemala",
		Category: "flag",
	},
	"🇬🇬": &emojiMeta{
		Keywords: []string{"gg", "flag", "nation", "country", "banner"},
		Name:     "guernsey",
		Category: "flag",
	},
	"🇬🇳": &emojiMeta{
		Keywords: []string{"gn", "flag", "nation", "country", "banner"},
		Name:     "guinea",
		Category: "flag",
	},
	"🇬🇼": &emojiMeta{
		Keywords: []string{"gw", "bissau", "flag", "nation", "country", "banner"},
		Name:     "guinea_bissau",
		Category: "flag",
	},
	"🇬🇾": &emojiMeta{
		Keywords: []string{"gy", "flag", "nation", "country", "banner"},
		Name:     "guyana",
		Category: "flag",
	},
	"🇭🇹": &emojiMeta{
		Keywords: []string{"ht", "flag", "nation", "country", "banner"},
		Name:     "haiti",
		Category: "flag",
	},
	"🇭🇳": &emojiMeta{
		Keywords: []string{"hn", "flag", "nation", "country", "banner"},
		Name:     "honduras",
		Category: "flag",
	},
	"🇭🇰": &emojiMeta{
		Keywords: []string{"hong", "kong", "flag", "nation", "country", "banner"},
		Name:     "hong_kong",
		Category: "flag",
	},
	"🇭🇺": &emojiMeta{
		Keywords: []string{"hu", "flag", "nation", "country", "banner"},
		Name:     "hungary",
		Category: "flag",
	},
	"🇮🇸": &emojiMeta{
		Keywords: []string{"is", "flag", "nation", "country", "banner"},
		Name:     "iceland",
		Category: "flag",
	},
	"🇮🇳": &emojiMeta{
		Keywords: []string{"in", "flag", "nation", "country", "banner"},
		Name:     "india",
		Category: "flag",
	},
	"🇮🇩": &emojiMeta{
		Keywords: []string{"flag", "nation", "country", "banner"},
		Name:     "indonesia",
		Category: "flag",
	},
	"🇮🇷": &emojiMeta{
		Keywords: []string{"iran,", "islamic", "republic", "flag", "nation", "country", "banner"},
		Name:     "iran",
		Category: "flag",
	},
	"🇮🇶": &emojiMeta{
		Keywords: []string{"iq", "flag", "nation", "country", "banner"},
		Name:     "iraq",
		Category: "flag",
	},
	"🇮🇪": &emojiMeta{
		Keywords: []string{"ie", "flag", "nation", "country", "banner"},
		Name:     "ireland",
		Category: "flag",
	},
	"🇮🇲": &emojiMeta{
		Keywords: []string{"isle", "man", "flag", "nation", "country", "banner"},
		Name:     "isle_of_man",
		Category: "flag",
	},
	"🇮🇱": &emojiMeta{
		Keywords: []string{"il", "flag", "nation", "country", "banner"},
		Name:     "israel",
		Category: "flag",
	},
	"🇮🇹": &emojiMeta{
		Keywords: []string{"italy", "flag", "nation", "country", "banner"},
		Name:     "it",
		Category: "flag",
	},
	"🇨🇮": &emojiMeta{
		Keywords: []string{"ivory", "coast", "flag", "nation", "country", "banner"},
		Name:     "cote_divoire",
		Category: "flag",
	},
	"🇯🇲": &emojiMeta{
		Keywords: []string{"jm", "flag", "nation", "country", "banner"},
		Name:     "jamaica",
		Category: "flag",
	},
	"🇯🇵": &emojiMeta{
		Keywords: []string{"japanese", "nation", "flag", "country", "banner"},
		Name:     "jp",
		Category: "flag",
	},
	"🇯🇪": &emojiMeta{
		Keywords: []string{"je", "flag", "nation", "country", "banner"},
		Name:     "jersey",
		Category: "flag",
	},
	"🇯🇴": &emojiMeta{
		Keywords: []string{"jo", "flag", "nation", "country", "banner"},
		Name:     "jordan",
		Category: "flag",
	},
	"🇰🇿": &emojiMeta{
		Keywords: []string{"kz", "flag", "nation", "country", "banner"},
		Name:     "kazakhstan",
		Category: "flag",
	},
	"🇰🇪": &emojiMeta{
		Keywords: []string{"ke", "flag", "nation", "country", "banner"},
		Name:     "kenya",
		Category: "flag",
	},
	"🇰🇮": &emojiMeta{
		Keywords: []string{"ki", "flag", "nation", "country", "banner"},
		Name:     "kiribati",
		Category: "flag",
	},
	"🇽🇰": &emojiMeta{
		Keywords: []string{"xk", "flag", "nation", "country", "banner"},
		Name:     "kosovo",
		Category: "flag",
	},
	"🇰🇼": &emojiMeta{
		Keywords: []string{"kw", "flag", "nation", "country", "banner"},
		Name:     "kuwait",
		Category: "flag",
	},
	"🇰🇬": &emojiMeta{
		Keywords: []string{"kg", "flag", "nation", "country", "banner"},
		Name:     "kyrgyzstan",
		Category: "flag",
	},
	"🇱🇦": &emojiMeta{
		Keywords: []string{"lao", "democratic", "republic", "flag", "nation", "country", "banner"},
		Name:     "laos",
		Category: "flag",
	},
	"🇱🇻": &emojiMeta{
		Keywords: []string{"lv", "flag", "nation", "country", "banner"},
		Name:     "latvia",
		Category: "flag",
	},
	"🇱🇧": &emojiMeta{
		Keywords: []string{"lb", "flag", "nation", "country", "banner"},
		Name:     "lebanon",
		Category: "flag",
	},
	"🇱🇸": &emojiMeta{
		Keywords: []string{"ls", "flag", "nation", "country", "banner"},
		Name:     "lesotho",
		Category: "flag",
	},
	"🇱🇷": &emojiMeta{
		Keywords: []string{"lr", "flag", "nation", "country", "banner"},
		Name:     "liberia",
		Category: "flag",
	},
	"🇱🇾": &emojiMeta{
		Keywords: []string{"ly", "flag", "nation", "country", "banner"},
		Name:     "libya",
		Category: "flag",
	},
	"🇱🇮": &emojiMeta{
		Keywords: []string{"li", "flag", "nation", "country", "banner"},
		Name:     "liechtenstein",
		Category: "flag",
	},
	"🇱🇹": &emojiMeta{
		Keywords: []string{"lt", "flag", "nation", "country", "banner"},
		Name:     "lithuania",
		Category: "flag",
	},
	"🇱🇺": &emojiMeta{
		Keywords: []string{"lu", "flag", "nation", "country", "banner"},
		Name:     "luxembourg",
		Category: "flag",
	},
	"🇲🇴": &emojiMeta{
		Keywords: []string{"macao", "flag", "nation", "country", "banner"},
		Name:     "macau",
		Category: "flag",
	},
	"🇲🇰": &emojiMeta{
		Keywords: []string{"macedonia,", "flag", "nation", "country", "banner"},
		Name:     "macedonia",
		Category: "flag",
	},
	"🇲🇬": &emojiMeta{
		Keywords: []string{"mg", "flag", "nation", "country", "banner"},
		Name:     "madagascar",
		Category: "flag",
	},
	"🇲🇼": &emojiMeta{
		Keywords: []string{"mw", "flag", "nation", "country", "banner"},
		Name:     "malawi",
		Category: "flag",
	},
	"🇲🇾": &emojiMeta{
		Keywords: []string{"my", "flag", "nation", "country", "banner"},
		Name:     "malaysia",
		Category: "flag",
	},
	"🇲🇻": &emojiMeta{
		Keywords: []string{"mv", "flag", "nation", "country", "banner"},
		Name:     "maldives",
		Category: "flag",
	},
	"🇲🇱": &emojiMeta{
		Keywords: []string{"ml", "flag", "nation", "country", "banner"},
		Name:     "mali",
		Category: "flag",
	},
	"🇲🇹": &emojiMeta{
		Keywords: []string{"mt", "flag", "nation", "country", "banner"},
		Name:     "malta",
		Category: "flag",
	},
	"🇲🇭": &emojiMeta{
		Keywords: []string{"marshall", "islands", "flag", "nation", "country", "banner"},
		Name:     "marshall_islands",
		Category: "flag",
	},
	"🇲🇶": &emojiMeta{
		Keywords: []string{"mq", "flag", "nation", "country", "banner"},
		Name:     "martinique",
		Category: "flag",
	},
	"🇲🇷": &emojiMeta{
		Keywords: []string{"mr", "flag", "nation", "country", "banner"},
		Name:     "mauritania",
		Category: "flag",
	},
	"🇲🇺": &emojiMeta{
		Keywords: []string{"mu", "flag", "nation", "country", "banner"},
		Name:     "mauritius",
		Category: "flag",
	},
	"🇾🇹": &emojiMeta{
		Keywords: []string{"yt", "flag", "nation", "country", "banner"},
		Name:     "mayotte",
		Category: "flag",
	},
	"🇲🇽": &emojiMeta{
		Keywords: []string{"mx", "flag", "nation", "country", "banner"},
		Name:     "mexico",
		Category: "flag",
	},
	"🇫🇲": &emojiMeta{
		Keywords: []string{"micronesia,", "federated", "states", "flag", "nation", "country", "banner"},
		Name:     "micronesia",
		Category: "flag",
	},
	"🇲🇩": &emojiMeta{
		Keywords: []string{"moldova,", "republic", "flag", "nation", "country", "banner"},
		Name:     "moldova",
		Category: "flag",
	},
	"🇲🇨": &emojiMeta{
		Keywords: []string{"mc", "flag", "nation", "country", "banner"},
		Name:     "monaco",
		Category: "flag",
	},
	"🇲🇳": &emojiMeta{
		Keywords: []string{"mn", "flag", "nation", "country", "banner"},
		Name:     "mongolia",
		Category: "flag",
	},
	"🇲🇪": &emojiMeta{
		Keywords: []string{"me", "flag", "nation", "country", "banner"},
		Name:     "montenegro",
		Category: "flag",
	},
	"🇲🇸": &emojiMeta{
		Keywords: []string{"ms", "flag", "nation", "country", "banner"},
		Name:     "montserrat",
		Category: "flag",
	},
	"🇲🇦": &emojiMeta{
		Keywords: []string{"ma", "flag", "nation", "country", "banner"},
		Name:     "morocco",
		Category: "flag",
	},
	"🇲🇿": &emojiMeta{
		Keywords: []string{"mz", "flag", "nation", "country", "banner"},
		Name:     "mozambique",
		Category: "flag",
	},
	"🇲🇲": &emojiMeta{
		Keywords: []string{"mm", "flag", "nation", "country", "banner"},
		Name:     "myanmar",
		Category: "flag",
	},
	"🇳🇦": &emojiMeta{
		Keywords: []string{"na", "flag", "nation", "country", "banner"},
		Name:     "namibia",
		Category: "flag",
	},
	"🇳🇷": &emojiMeta{
		Keywords: []string{"nr", "flag", "nation", "country", "banner"},
		Name:     "nauru",
		Category: "flag",
	},
	"🇳🇵": &emojiMeta{
		Keywords: []string{"np", "flag", "nation", "country", "banner"},
		Name:     "nepal",
		Category: "flag",
	},
	"🇳🇱": &emojiMeta{
		Keywords: []string{"nl", "flag", "nation", "country", "banner"},
		Name:     "netherlands",
		Category: "flag",
	},
	"🇳🇨": &emojiMeta{
		Keywords: []string{"new", "caledonia", "flag", "nation", "country", "banner"},
		Name:     "new_caledonia",
		Category: "flag",
	},
	"🇳🇿": &emojiMeta{
		Keywords: []string{"new", "zealand", "flag", "nation", "country", "banner"},
		Name:     "new_zealand",
		Category: "flag",
	},
	"🇳🇮": &emojiMeta{
		Keywords: []string{"ni", "flag", "nation", "country", "banner"},
		Name:     "nicaragua",
		Category: "flag",
	},
	"🇳🇪": &emojiMeta{
		Keywords: []string{"ne", "flag", "nation", "country", "banner"},
		Name:     "niger",
		Category: "flag",
	},
	"🇳🇬": &emojiMeta{
		Keywords: []string{"flag", "nation", "country", "banner"},
		Name:     "nigeria",
		Category: "flag",
	},
	"🇳🇺": &emojiMeta{
		Keywords: []string{"nu", "flag", "nation", "country", "banner"},
		Name:     "niue",
		Category: "flag",
	},
	"🇳🇫": &emojiMeta{
		Keywords: []string{"norfolk", "island", "flag", "nation", "country", "banner"},
		Name:     "norfolk_island",
		Category: "flag",
	},
	"🇲🇵": &emojiMeta{
		Keywords: []string{"northern", "mariana", "islands", "flag", "nation", "country", "banner"},
		Name:     "northern_mariana_islands",
		Category: "flag",
	},
	"🇰🇵": &emojiMeta{
		Keywords: []string{"north", "korea", "nation", "flag", "country", "banner"},
		Name:     "north_korea",
		Category: "flag",
	},
	"🇳🇴": &emojiMeta{
		Keywords: []string{"no", "flag", "nation", "country", "banner"},
		Name:     "norway",
		Category: "flag",
	},
	"🇴🇲": &emojiMeta{
		Keywords: []string{"om_symbol", "flag", "nation", "country", "banner"},
		Name:     "oman",
		Category: "flag",
	},
	"🇵🇰": &emojiMeta{
		Keywords: []string{"pk", "flag", "nation", "country", "banner"},
		Name:     "pakistan",
		Category: "flag",
	},
	"🇵🇼": &emojiMeta{
		Keywords: []string{"pw", "flag", "nation", "country", "banner"},
		Name:     "palau",
		Category: "flag",
	},
	"🇵🇸": &emojiMeta{
		Keywords: []string{"palestine", "palestinian", "territories", "flag", "nation", "country", "banner"},
		Name:     "palestinian_territories",
		Category: "flag",
	},
	"🇵🇦": &emojiMeta{
		Keywords: []string{"pa", "flag", "nation", "country", "banner"},
		Name:     "panama",
		Category: "flag",
	},
	"🇵🇬": &emojiMeta{
		Keywords: []string{"papua", "new", "guinea", "flag", "nation", "country", "banner"},
		Name:     "papua_new_guinea",
		Category: "flag",
	},
	"🇵🇾": &emojiMeta{
		Keywords: []string{"py", "flag", "nation", "country", "banner"},
		Name:     "paraguay",
		Category: "flag",
	},
	"🇵🇪": &emojiMeta{
		Keywords: []string{"pe", "flag", "nation", "country", "banner"},
		Name:     "peru",
		Category: "flag",
	},
	"🇵🇭": &emojiMeta{
		Keywords: []string{"ph", "flag", "nation", "country", "banner"},
		Name:     "philippines",
		Category: "flag",
	},
	"🇵🇳": &emojiMeta{
		Keywords: []string{"pitcairn", "flag", "nation", "country", "banner"},
		Name:     "pitcairn_islands",
		Category: "flag",
	},
	"🇵🇱": &emojiMeta{
		Keywords: []string{"pl", "flag", "nation", "country", "banner"},
		Name:     "poland",
		Category: "flag",
	},
	"🇵🇹": &emojiMeta{
		Keywords: []string{"pt", "flag", "nation", "country", "banner"},
		Name:     "portugal",
		Category: "flag",
	},
	"🇵🇷": &emojiMeta{
		Keywords: []string{"puerto", "rico", "flag", "nation", "country", "banner"},
		Name:     "puerto_rico",
		Category: "flag",
	},
	"🇶🇦": &emojiMeta{
		Keywords: []string{"qa", "flag", "nation", "country", "banner"},
		Name:     "qatar",
		Category: "flag",
	},
	"🇷🇪": &emojiMeta{
		Keywords: []string{"réunion", "flag", "nation", "country", "banner"},
		Name:     "reunion",
		Category: "flag",
	},
	"🇷🇴": &emojiMeta{
		Keywords: []string{"ro", "flag", "nation", "country", "banner"},
		Name:     "romania",
		Category: "flag",
	},
	"🇷🇺": &emojiMeta{
		Keywords: []string{"russian", "federation", "flag", "nation", "country", "banner"},
		Name:     "ru",
		Category: "flag",
	},
	"🇷🇼": &emojiMeta{
		Keywords: []string{"rw", "flag", "nation", "country", "banner"},
		Name:     "rwanda",
		Category: "flag",
	},
	"🇧🇱": &emojiMeta{
		Keywords: []string{"saint", "barthélemy", "flag", "nation", "country", "banner"},
		Name:     "st_barthelemy",
		Category: "flag",
	},
	"🇸🇭": &emojiMeta{
		Keywords: []string{"saint", "helena", "ascension", "tristan", "cunha", "flag", "nation", "country", "banner"},
		Name:     "st_helena",
		Category: "flag",
	},
	"🇰🇳": &emojiMeta{
		Keywords: []string{"saint", "kitts", "nevis", "flag", "nation", "country", "banner"},
		Name:     "st_kitts_nevis",
		Category: "flag",
	},
	"🇱🇨": &emojiMeta{
		Keywords: []string{"saint", "lucia", "flag", "nation", "country", "banner"},
		Name:     "st_lucia",
		Category: "flag",
	},
	"🇵🇲": &emojiMeta{
		Keywords: []string{"saint", "pierre", "miquelon", "flag", "nation", "country", "banner"},
		Name:     "st_pierre_miquelon",
		Category: "flag",
	},
	"🇻🇨": &emojiMeta{
		Keywords: []string{"saint", "vincent", "grenadines", "flag", "nation", "country", "banner"},
		Name:     "st_vincent_grenadines",
		Category: "flag",
	},
	"🇼🇸": &emojiMeta{
		Keywords: []string{"ws", "flag", "nation", "country", "banner"},
		Name:     "samoa",
		Category: "flag",
	},
	"🇸🇲": &emojiMeta{
		Keywords: []string{"san", "marino", "flag", "nation", "country", "banner"},
		Name:     "san_marino",
		Category: "flag",
	},
	"🇸🇹": &emojiMeta{
		Keywords: []string{"sao", "tome", "principe", "flag", "nation", "country", "banner"},
		Name:     "sao_tome_principe",
		Category: "flag",
	},
	"🇸🇦": &emojiMeta{
		Keywords: []string{"flag", "nation", "country", "banner"},
		Name:     "saudi_arabia",
		Category: "flag",
	},
	"🇸🇳": &emojiMeta{
		Keywords: []string{"sn", "flag", "nation", "country", "banner"},
		Name:     "senegal",
		Category: "flag",
	},
	"🇷🇸": &emojiMeta{
		Keywords: []string{"rs", "flag", "nation", "country", "banner"},
		Name:     "serbia",
		Category: "flag",
	},
	"🇸🇨": &emojiMeta{
		Keywords: []string{"sc", "flag", "nation", "country", "banner"},
		Name:     "seychelles",
		Category: "flag",
	},
	"🇸🇱": &emojiMeta{
		Keywords: []string{"sierra", "leone", "flag", "nation", "country", "banner"},
		Name:     "sierra_leone",
		Category: "flag",
	},
	"🇸🇬": &emojiMeta{
		Keywords: []string{"sg", "flag", "nation", "country", "banner"},
		Name:     "singapore",
		Category: "flag",
	},
	"🇸🇽": &emojiMeta{
		Keywords: []string{"sint", "maarten", "dutch", "flag", "nation", "country", "banner"},
		Name:     "sint_maarten",
		Category: "flag",
	},
	"🇸🇰": &emojiMeta{
		Keywords: []string{"sk", "flag", "nation", "country", "banner"},
		Name:     "slovakia",
		Category: "flag",
	},
	"🇸🇮": &emojiMeta{
		Keywords: []string{"si", "flag", "nation", "country", "banner"},
		Name:     "slovenia",
		Category: "flag",
	},
	"🇸🇧": &emojiMeta{
		Keywords: []string{"solomon", "islands", "flag", "nation", "country", "banner"},
		Name:     "solomon_islands",
		Category: "flag",
	},
	"🇸🇴": &emojiMeta{
		Keywords: []string{"so", "flag", "nation", "country", "banner"},
		Name:     "somalia",
		Category: "flag",
	},
	"🇿🇦": &emojiMeta{
		Keywords: []string{"south", "africa", "flag", "nation", "country", "banner"},
		Name:     "south_africa",
		Category: "flag",
	},
	"🇬🇸": &emojiMeta{
		Keywords: []string{"south", "georgia", "sandwich", "islands", "flag", "nation", "country", "banner"},
		Name:     "south_georgia_south_sandwich_islands",
		Category: "flag",
	},
	"🇰🇷": &emojiMeta{
		Keywords: []string{"south", "korea", "nation", "flag", "country", "banner"},
		Name:     "kr",
		Category: "flag",
	},
	"🇸🇸": &emojiMeta{
		Keywords: []string{"south", "sd", "flag", "nation", "country", "banner"},
		Name:     "south_sudan",
		Category: "flag",
	},
	"🇪🇸": &emojiMeta{
		Keywords: []string{"spain", "flag", "nation", "country", "banner"},
		Name:     "es",
		Category: "flag",
	},
	"🇱🇰": &emojiMeta{
		Keywords: []string{"sri", "lanka", "flag", "nation", "country", "banner"},
		Name:     "sri_lanka",
		Category: "flag",
	},
	"🇸🇩": &emojiMeta{
		Keywords: []string{"sd", "flag", "nation", "country", "banner"},
		Name:     "sudan",
		Category: "flag",
	},
	"🇸🇷": &emojiMeta{
		Keywords: []string{"sr", "flag", "nation", "country", "banner"},
		Name:     "suriname",
		Category: "flag",
	},
	"🇸🇿": &emojiMeta{
		Keywords: []string{"sz", "flag", "nation", "country", "banner"},
		Name:     "swaziland",
		Category: "flag",
	},
	"🇸🇪": &emojiMeta{
		Keywords: []string{"se", "flag", "nation", "country", "banner"},
		Name:     "sweden",
		Category: "flag",
	},
	"🇨🇭": &emojiMeta{
		Keywords: []string{"ch", "flag", "nation", "country", "banner"},
		Name:     "switzerland",
		Category: "flag",
	},
	"🇸🇾": &emojiMeta{
		Keywords: []string{"syrian", "arab", "republic", "flag", "nation", "country", "banner"},
		Name:     "syria",
		Category: "flag",
	},
	"🇹🇼": &emojiMeta{
		Keywords: []string{"tw", "flag", "nation", "country", "banner"},
		Name:     "taiwan",
		Category: "flag",
	},
	"🇹🇯": &emojiMeta{
		Keywords: []string{"tj", "flag", "nation", "country", "banner"},
		Name:     "tajikistan",
		Category: "flag",
	},
	"🇹🇿": &emojiMeta{
		Keywords: []string{"tanzania,", "united", "republic", "flag", "nation", "country", "banner"},
		Name:     "tanzania",
		Category: "flag",
	},
	"🇹🇭": &emojiMeta{
		Keywords: []string{"th", "flag", "nation", "country", "banner"},
		Name:     "thailand",
		Category: "flag",
	},
	"🇹🇱": &emojiMeta{
		Keywords: []string{"timor", "leste", "flag", "nation", "country", "banner"},
		Name:     "timor_leste",
		Category: "flag",
	},
	"🇹🇬": &emojiMeta{
		Keywords: []string{"tg", "flag", "nation", "country", "banner"},
		Name:     "togo",
		Category: "flag",
	},
	"🇹🇰": &emojiMeta{
		Keywords: []string{"tk", "flag", "nation", "country", "banner"},
		Name:     "tokelau",
		Category: "flag",
	},
	"🇹🇴": &emojiMeta{
		Keywords: []string{"to", "flag", "nation", "country", "banner"},
		Name:     "tonga",
		Category: "flag",
	},
	"🇹🇹": &emojiMeta{
		Keywords: []string{"trinidad", "tobago", "flag", "nation", "country", "banner"},
		Name:     "trinidad_tobago",
		Category: "flag",
	},
	"🇹🇳": &emojiMeta{
		Keywords: []string{"tn", "flag", "nation", "country", "banner"},
		Name:     "tunisia",
		Category: "flag",
	},
	"🇹🇷": &emojiMeta{
		Keywords: []string{"turkey", "flag", "nation", "country", "banner"},
		Name:     "tr",
		Category: "flag",
	},
	"🇹🇲": &emojiMeta{
		Keywords: []string{"flag", "nation", "country", "banner"},
		Name:     "turkmenistan",
		Category: "flag",
	},
	"🇹🇨": &emojiMeta{
		Keywords: []string{"turks", "caicos", "islands", "flag", "nation", "country", "banner"},
		Name:     "turks_caicos_islands",
		Category: "flag",
	},
	"🇹🇻": &emojiMeta{
		Keywords: []string{"flag", "nation", "country", "banner"},
		Name:     "tuvalu",
		Category: "flag",
	},
	"🇺🇬": &emojiMeta{
		Keywords: []string{"ug", "flag", "nation", "country", "banner"},
		Name:     "uganda",
		Category: "flag",
	},
	"🇺🇦": &emojiMeta{
		Keywords: []string{"ua", "flag", "nation", "country", "banner"},
		Name:     "ukraine",
		Category: "flag",
	},
	"🇦🇪": &emojiMeta{
		Keywords: []string{"united", "arab", "emirates", "flag", "nation", "country", "banner"},
		Name:     "united_arab_emirates",
		Category: "flag",
	},
	"🇬🇧": &emojiMeta{
		Keywords: []string{"united", "kingdom", "great", "britain", "northern", "ireland", "flag", "nation", "country", "banner", "british", "UK", "english", "england", "union jack"},
		Name:     "uk",
		Category: "flag",
	},
	"🏴󠁧󠁢󠁥󠁮󠁧󠁿": &emojiMeta{
		Keywords: []string{"flag", "english"},
		Name:     "england",
		Category: "flag",
	},
	"🏴󠁧󠁢󠁳󠁣󠁴󠁿": &emojiMeta{
		Keywords: []string{"flag", "scottish"},
		Name:     "scotland",
		Category: "flag",
	},
	"🏴󠁧󠁢󠁷󠁬󠁳󠁿": &emojiMeta{
		Keywords: []string{"flag", "welsh"},
		Name:     "wales",
		Category: "flag",
	},
	"🇺🇸": &emojiMeta{
		Keywords: []string{"united", "states", "america", "flag", "nation", "country", "banner"},
		Name:     "us",
		Category: "flag",
	},
	"🇻🇮": &emojiMeta{
		Keywords: []string{"virgin", "islands", "us", "flag", "nation", "country", "banner"},
		Name:     "us_virgin_islands",
		Category: "flag",
	},
	"🇺🇾": &emojiMeta{
		Keywords: []string{"uy", "flag", "nation", "country", "banner"},
		Name:     "uruguay",
		Category: "flag",
	},
	"🇺🇿": &emojiMeta{
		Keywords: []string{"uz", "flag", "nation", "country", "banner"},
		Name:     "uzbekistan",
		Category: "flag",
	},
	"🇻🇺": &emojiMeta{
		Keywords: []string{"vu", "flag", "nation", "country", "banner"},
		Name:     "vanuatu",
		Category: "flag",
	},
	"🇻🇦": &emojiMeta{
		Keywords: []string{"vatican", "city", "flag", "nation", "country", "banner"},
		Name:     "vatican_city",
		Category: "flag",
	},
	"🇻🇪": &emojiMeta{
		Keywords: []string{"ve", "bolivarian", "republic", "flag", "nation", "country", "banner"},
		Name:     "venezuela",
		Category: "flag",
	},
	"🇻🇳": &emojiMeta{
		Keywords: []string{"viet", "nam", "flag", "nation", "country", "banner"},
		Name:     "vietnam",
		Category: "flag",
	},
	"🇼🇫": &emojiMeta{
		Keywords: []string{"wallis", "futuna", "flag", "nation", "country", "banner"},
		Name:     "wallis_futuna",
		Category: "flag",
	},
	"🇪🇭": &emojiMeta{
		Keywords: []string{"western", "sahara", "flag", "nation", "country", "banner"},
		Name:     "western_sahara",
		Category: "flag",
	},
	"🇾🇪": &emojiMeta{
		Keywords: []string{"ye", "flag", "nation", "country", "banner"},
		Name:     "yemen",
		Category: "flag",
	},
	"🇿🇲": &emojiMeta{
		Keywords: []string{"zm", "flag", "nation", "country", "banner"},
		Name:     "zambia",
		Category: "flag",
	},
	"🇿🇼": &emojiMeta{
		Keywords: []string{"zw", "flag", "nation", "country", "banner"},
		Name:     "zimbabwe",
		Category: "flag",
	},
	"🇺🇳": &emojiMeta{
		Keywords: []string{"un", "flag", "banner"},
		Name:     "united_nations",
		Category: "flag",
	},
	"🏴‍☠️": &emojiMeta{
		Keywords: []string{"skull", "crossbones", "flag", "banner"},
		Name:     "pirate_flag",
		Category: "flag",
	},
}
