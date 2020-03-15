package maps

import "testing"

/* func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "就是个测试"}

	got := Search(dictionary, "test")
	want := "就是个测试"

	assertStrings(t, got, want)

} */

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "就是个测试"}

	t.Run("known world", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "就是个测试"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		def := "一个测试"

		err := dictionary.Add(word, def)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, def)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		def := "一个测试"
		dic := Dictionary{word: def}
		err := dic.Add(word, "新测试")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dic, word, def)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("已存在的字段", func(t *testing.T) {
		word := "test"
		def := "一个测试"
		newDef := "new definition"

		dic := Dictionary{word: def}
		err := dic.Update(word, newDef)

		assertError(t, err, nil)
		assertDefinition(t, dic, word, def)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		def := "一个测试"
		dic := Dictionary{}

		err := dic.Update(word, def)
		assertError(t, err, ErrWordDoesNotExist)
	})

}

func TestDelete(t *testing.T) {
	word := "test"
	dic := Dictionary{word: "测试定义"}

	dic.Delete(word)
	_, err := dic.Search(word)

	if err != ErrNotFound {
		t.Errorf("Expected '%s' to be deleted", word)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, def string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if def != got {
		t.Errorf("got '%s' want '%s'", got, def)
	}

}
