package bantuan

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloTable(b *testing.B) {
	benchmarks := []struct {
		nameFunc string
		request  string
	}{
		{
			nameFunc: "Habie",
			request:  "Habie Putra",
		},
		{
			nameFunc: "Norstein",
			request:  "Thomas H. Norstein",
		},
	}

	for _, bm := range benchmarks {
		b.Run(bm.nameFunc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloIsekai(bm.request)
			}
		})
	}

}

func BenchmarkHelloIsekai(b *testing.B) {
	b.Run("Norstein", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloIsekai("Norstein")
		}
	})

	b.Run("Eddy", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloIsekai("Eddy")
		}
	})

	b.Run("Tama", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloIsekai("Tama")
		}
	})

}

func TestSubTest(t *testing.T) {
	t.Run("Radit", func(t *testing.T) {
		hasil := HelloIsekai("Radit")
		assert.Equal(t, "Hi Radit", hasil, "Seharusnya Hi Radit")
	})

	t.Run("Awang", func(t *testing.T) {
		hasil := HelloIsekai("Awang")
		assert.Equal(t, "Hello from isekai, Awang", hasil, "Seharusnya Hello from isekai, Awang")
	})
}

func TestMain(m *testing.M) {
	fmt.Println("Koneksi ke Database")

	m.Run()

	fmt.Println("Selesai")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Unit test tidak bisa jalan di MAC OS")
	}

	hasil := HelloIsekai("Habie")
	assert.Equal(t, "Hello from isekai, Habie", hasil, "Salah coba cek lagi string expectednya")
	fmt.Println("Unit test dengan skip selesai")
}

func TestHelloHabieAssert(t *testing.T) {
	hasil := HelloIsekai("Habie")
	assert.Equal(t, "Hello from isekai, Habie", hasil, "Seharusnya output nya adalah Hello from isekai, Habie")
	fmt.Println("Test Menggunakan Assert Done")
}

func TestHelloHabieRequire(t *testing.T) {
	hasil := HelloIsekai("Habie")
	require.Equal(t, "Hi from isekai, Habie", hasil, "Seharusnya Hi from isekai, Habie")
	fmt.Println("Test Menggunakan Require Selesai") //ini jika terjadi error tidak akan dijalankan
}

func TestHellooHabie(t *testing.T) {
	hasil := HelloIsekai("Habie")
	if hasil != "Hello from isekai, Habie" {
		//error
		//panc("Hasil bukan 'Hallo Habie'")
		//t.Fail()
		//menangkap error dan melanjutkan program
		t.Error("Seharusnya output nya adalah Hello from isekai, Habie")
	}

	fmt.Println("Testing selesai")
}

func TestHellooNorstein(t *testing.T) {
	hasil := HelloIsekai("Norstein")
	if hasil != "Hello from isekai, Norstein" {
		//error
		//panc("Hasil bukan 'Hallo Habie'")
		//t.Fail()
		//t.fatal lansung menghentikan program tanpa melanjutkan kode yang ada dibawahnya
		t.Fatal("Seharusnya output nya adalah Hello from isekai, Norstein")
	}
	fmt.Println("Testing berhasil")
}

func TestTableTest(t *testing.T) {
	data := []struct {
		name     string
		request  string
		expected string
	}{{
		name:     "Goku",
		request:  "Goku",
		expected: "Hello from isekai, Goku",
	}, {
		name:     "Kazuya",
		request:  "Kazuya",
		expected: "Hello from isekai, Kazuya"},
	}

	for _, test := range data {
		t.Run(test.name, func(t *testing.T) {
			hasil := HelloIsekai(test.request)
			assert.Equal(t, test.expected, hasil, "Salah distring")
		})
	}
}
