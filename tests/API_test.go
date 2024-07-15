package tests

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
	docsv1 "github.com/wlcmtunknwndth/vk_test/proto/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strconv"
	"testing"
	"time"
)

const Sec = uint64(10_000_000_000)

var addr = "localhost:8888"

func TestCreate(t *testing.T) {
	conn, err := ConnectToServer()
	require.NoError(t, err)

	ctx := context.Background()
	tcUsr, tc := GenerateCases(15)

	delta := Sec // 1s in nanosecs

	for i, _ := range tc {
		tdoc, err := conn.Create(ctx, &tcUsr[i], grpc.StaticMethod())
		require.NoError(t, err)

		if tdoc.FetchTime-tc[i].FetchTime > delta || tdoc.FirstFetchTime-tc[i].FirstFetchTime > delta || tdoc.PubDate-tc[i].PubDate > delta {
			t.Error("big delay", tdoc.FetchTime-tc[i].FetchTime, tdoc.FirstFetchTime-tc[i].FirstFetchTime, tdoc.PubDate-tc[i].PubDate)
			return
		}
	}
}

func TestUpdate(t *testing.T) {
	conn, err := ConnectToServer()
	require.NoError(t, err)

	ctx := context.Background()
	tcUsr1, err := GenerateTDocsUsr(10)
	require.NoError(t, err)

	tcUsr2, err := GenerateTDocsUsr(10)
	require.NoError(t, err)
	// save
	//var tc1 = make([]docsv1.TDocument, 0, 10)
	for i, _ := range tcUsr1 {
		tc1, err := conn.Create(ctx, &tcUsr1[i], grpc.StaticMethod())
		require.NoError(t, err)

		tcUsr2[i].Url = tcUsr1[i].Url

		tdoc, err := conn.Update(ctx, &tcUsr2[i], grpc.StaticMethod())
		require.NoError(t, err)

		if tdoc.Text != tcUsr2[i].Text {
			t.Error("Text", tdoc.Text, "!=", tcUsr2[i].Text)
			return
		}
		if uint64(time.Now().UnixNano())-tc1.FetchTime > Sec {
			t.Error("Fetch Time", tdoc.FetchTime, "> 1s ago")
			return
		}
		if tdoc.FirstFetchTime != tc1.FirstFetchTime {
			t.Error("FirstFetchTime", tdoc.FirstFetchTime, "!=", tc1.FirstFetchTime)
			return
		}
		if tdoc.PubDate != tc1.PubDate {
			t.Error("PubDate", tdoc.PubDate, "!=", tc1.PubDate)
			return
		}

	}
}

func TestProcess(t *testing.T) {
	conn, err := ConnectToServer()
	require.NoError(t, err)

	tc, err := GenerateTDocs(10)
	require.NoError(t, err)

	ctx := context.Background()
	for i, _ := range tc {
		i := i
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			tc[i].FirstFetchTime = 0
			tdoc, err := conn.Process(ctx, &tc[i], grpc.StaticMethod())
			require.NoError(t, err)

			now := uint64(time.Now().UnixNano())
			if now-tdoc.FirstFetchTime > Sec {
				t.Error("FirstFetchTime:", now, "-", tdoc.FirstFetchTime, "> 1s")
				return
			}

			if now-tdoc.FetchTime > Sec {
				t.Error("FetchTime:", now, "-", tdoc.FetchTime, "> 1s")
				return
			}

			if tdoc.Url != tc[i].Url {
				t.Error("different texts")
			}

			if tdoc.PubDate != tc[i].PubDate {
				t.Error("PubDate:", now, "!=", tdoc.PubDate)
				return
			}
		})
	}
}

func ConnectToServer() (docsv1.DocumentsClient, error) {
	cl, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return docsv1.NewDocumentsClient(cl), nil
}

func GenerateCases(n int) ([]docsv1.UserTDocument, []docsv1.TDocument) {

	//lis := bufconn.Listen(1024 * 1024)
	//conn, err := grpc.NewClient("tcp", grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
	//	return lis.Dial()
	//}), grpc.WithTransportCredentials(insecure.NewCredentials()))

	testCasesUsr := make([]docsv1.UserTDocument, n)
	testCases := make([]docsv1.TDocument, n)

	for i := 0; i < n; i++ {
		msg, err := gofakeit.EmailText(&gofakeit.EmailOptions{})
		url := gofakeit.URL()

		//pubDate := gofakeit.Uint64()
		//fetchTime := gofakeit.Uint64()
		//firstFetchTime := gofakeit.Uint64()

		now := uint64(time.Now().UnixNano())
		if err != nil {
			return testCasesUsr, testCases
		}
		testCasesUsr[i] = docsv1.UserTDocument{
			Url:  url,
			Text: msg,
		}
		testCases[i] = docsv1.TDocument{
			Url:            url,
			PubDate:        now,
			FetchTime:      now,
			Text:           msg,
			FirstFetchTime: now,
		}
	}

	return testCasesUsr, testCases
}

func GenerateTDocsUsr(n int) ([]docsv1.UserTDocument, error) {
	var testCasesUsr = make([]docsv1.UserTDocument, n)
	for i := 0; i < n; i++ {
		msg, err := gofakeit.EmailText(&gofakeit.EmailOptions{})
		if err != nil {
			return nil, err
		}
		url := gofakeit.URL()

		//pubDate := gofakeit.Uint64()
		//fetchTime := gofakeit.Uint64()
		//firstFetchTime := gofakeit.Uint64()

		testCasesUsr[i] = docsv1.UserTDocument{
			Url:  url,
			Text: msg,
		}
	}

	return testCasesUsr, nil
}

func GenerateTDocs(n int) ([]docsv1.TDocument, error) {
	var testCases = make([]docsv1.TDocument, n)
	for i := 0; i < n; i++ {
		msg, err := gofakeit.EmailText(&gofakeit.EmailOptions{})
		if err != nil {
			return nil, err
		}
		url := gofakeit.URL()
		now := uint64(time.Now().UnixNano())
		//pubDate := gofakeit.Uint64()
		//fetchTime := gofakeit.Uint64()
		//firstFetchTime := gofakeit.Uint64()

		testCases[i] = docsv1.TDocument{
			Url:            url,
			PubDate:        now,
			FetchTime:      now,
			Text:           msg,
			FirstFetchTime: now,
		}
	}

	return testCases, nil
}
